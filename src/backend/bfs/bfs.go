package bfs

import (
	"container/list"
	"fmt"
	"net/http"
	// "golang.org/x/net/html"
	"encoding/json"
	"strings"
	"time"
	"github.com/gocolly/colly"
	"sync"
)

type Link struct {
	URL   string
}

type Request struct {
    Start  string `json:"start"`
    Target string `json:"target"`
}

var linkCache = &sync.Map{}

// fungsi getLinks yang pake go-colly
func getLinks(pageTitle string) []Link {
    pageURL := "https://en.wikipedia.org/wiki/" + pageTitle

    if links, ok := linkCache.Load(pageURL); ok {
        return links.([]Link)
    }

    c := colly.NewCollector()
    var links []Link

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") && !strings.Contains(link, "category") && !strings.Contains(link, "Main_Page") && link != "/wiki/Wikipedia:About" && link != "/wiki/Wikipedia:General_disclaimer" && link != "/wiki/Wikipedia:Contact_us" && link != "/wiki/Special:SpecialPages" {
            isArticleLink := true
            for _, class := range strings.Fields(e.Attr("class")) {
                if class == "new" || strings.Contains(strings.ToLower(class), "portal") {
                    isArticleLink = false
                    break
                }
            }
            if isArticleLink && !strings.Contains(link, ":") {
                // Extract the page title from the URL
                pageTitle := strings.TrimPrefix(link, "/wiki/")
                links = append(links, Link{URL: pageTitle})
            }
        }
    })

    c.Visit(pageURL)

    linkCache.Store(pageURL, links)

    return links
}

func findShortestPath(startPage, endPage string) ([]Link, time.Duration) {
    startTime := time.Now()
    queue := list.New()
    visited := make(map[string]bool)
    path := make(map[string][]Link)
    queue.PushBack([]Link{{URL: startPage}})

    for queue.Len() > 0 {
        currentPath := queue.Remove(queue.Front()).([]Link)
        currentLink := currentPath[len(currentPath)-1]

        if currentLink.URL == endPage {
            return currentPath, time.Since(startTime)
        }

        links := getLinks(currentLink.URL)
        for _, link := range links {
            if !visited[link.URL] {
                visited[link.URL] = true
                newPath := append(currentPath, link)
                queue.PushBack(newPath)
                path[link.URL] = newPath
                fmt.Print(newPath, "\n")

                if link.URL == endPage {
                    return newPath, time.Since(startTime)
                }
            }
        }
    }

    return nil, time.Since(startTime)
}

func BFSHandler(w http.ResponseWriter, r *http.Request) {
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

// func main() {
// 	// startTime := time.Now()
// 	// startURL := "https://en.wikipedia.org/wiki/Joko_Widodo"
// 	// endURL := "https://en.wikipedia.org/wiki/Koi"

// 	// shortestPath := findShortestPath(startURL, endURL)
// 	// if shortestPath == nil {
// 	// 	log.Fatal("No path found")
// 	// }

// 	// fmt.Println("Shortest path:")
// 	// for _, link := range shortestPath {
// 	// 	fmt.Println(link.URL)
// 	// }
// 	// endTime := time.Now()
// 	// elapsed := endTime.Sub(startTime)
// 	// fmt.Println("Execution time:", elapsed)

// 	mux := http.NewServeMux()
//     mux.HandleFunc("/bfs", handler)

//     // Setup CORS
//     c := cors.New(cors.Options{
//         AllowedOrigins: []string{"*"},
//         AllowCredentials: true,
//         AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
//     })

//     handler := c.Handler(mux)

//     http.ListenAndServe(":8080", handler)
// }
