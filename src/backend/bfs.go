package main

import (
    "container/list"
    "fmt"
    "net/http"
    "golang.org/x/net/html"
	"encoding/json"
	"github.com/rs/cors"
    "strings"
	"time"
)

// Link merepresentasikan tautan antara halaman Wikipedia
type Link struct {
    URL   string
}

// Fungsi getLinks mengambil tautan dari halaman Wikipedia
func getLinks(pageURL string) []Link {
	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	var links []Link
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, "/wiki/") {
					// Periksa apakah tautan memiliki class "new" atau title yang mengandung kata kunci "Portal"
					isArticleLink := true
					for _, class := range strings.Fields(attr.Val) {
						if class == "new" || strings.Contains(strings.ToLower(class), "portal") {
							isArticleLink = false
							break
						}
					}
					// Periksa apakah tautan mengandung pola khas untuk URL halaman Wikipedia
					if isArticleLink && strings.HasPrefix(attr.Val, "/wiki/") && !strings.Contains(attr.Val, ":") {
						link := Link{
							URL:   "https://en.wikipedia.org" + attr.Val,
						}
						links = append(links, link)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return links
}

// Algoritma BFS untuk menemukan jalur terpendek antara dua halaman Wikipedia
func findShortestPath(startURL, endURL string) ([]Link, time.Duration) {
	startTime := time.Now()
	// Queue untuk BFS
	queue := list.New()

	// Menyimpan jalur yang telah dikunjungi
	visited := make(map[string]bool)

	// Menyimpan jalur yang telah ditemukan
	path := make(map[string][]Link)

	// Tambahkan halaman awal ke queue
	queue.PushBack([]Link{{URL: startURL}})

	// Lakukan BFS
	for queue.Len() > 0 {
		// Ambil jalur dari queue
		currentPath := queue.Remove(queue.Front()).([]Link)
		currentLink := currentPath[len(currentPath)-1]

		// Jika sudah mencapai halaman akhir
		if currentLink.URL == endURL {
			return currentPath, time.Since(startTime)
		}

		// Periksa halaman yang terhubung dengan halaman saat ini
		links := getLinks(currentLink.URL)
		for _, link := range links {
			// Jika halaman belum dikunjungi
			if !visited[link.URL] {
				// Tandai halaman sebagai dikunjungi
				visited[link.URL] = true

				// Tambahkan link ke jalur yang sedang diperiksa
				newPath := append(currentPath, link)

				// Tambahkan jalur ke queue
				queue.PushBack(newPath)

				// Simpan jalur
				path[link.URL] = newPath
				fmt.Print(newPath, "\n")

				// Jika link adalah endURL, langsung kembalikan jalur yang ditemukan
				if link.URL == endURL {
					return newPath, time.Since(startTime)
				}
			}
		}
	}

	// Jika tidak ada jalur yang ditemukan
	return nil, time.Since(startTime)
}

func handler(w http.ResponseWriter, r *http.Request) {
    start := r.URL.Query().Get("start")
    target := r.URL.Query().Get("target")

    if start == "" || target == "" {
        http.Error(w, "Missing start or target parameter", http.StatusBadRequest)
        return
    }

    shortestPath, duration := findShortestPath(start, target)

    if shortestPath != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "path": shortestPath,
            "duration": duration,
        })
    } else {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "error": "No path found",
        })
    }
}

func main() {
	// startTime := time.Now()
	// startURL := "https://en.wikipedia.org/wiki/Joko_Widodo"
	// endURL := "https://en.wikipedia.org/wiki/Koi"

	// shortestPath := findShortestPath(startURL, endURL)
	// if shortestPath == nil {
	// 	log.Fatal("No path found")
	// }

	// fmt.Println("Shortest path:")
	// for _, link := range shortestPath {
	// 	fmt.Println(link.URL)
	// }
	// endTime := time.Now()
	// elapsed := endTime.Sub(startTime)
	// fmt.Println("Execution time:", elapsed)

	mux := http.NewServeMux()
    mux.HandleFunc("/bfs", handler)

    // Setup CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    })

    handler := c.Handler(mux)

    http.ListenAndServe(":8080", handler)
}
