package main

import (
    "fmt"
    "github.com/gocolly/colly"
    "github.com/rs/cors"
    "strings"
	"time"
    "sync"
    "net/http"
    "encoding/json"
)

// Link merepresentasikan tautan antara halaman Wikipedia
type Link struct {
    URL   string
}

type Request struct {
    Start  string `json:"start"`
    Target string `json:"target"`
}

var linkCache = &sync.Map{}

// fungsi getLinks yang pake go-colly
func getLinks(pageURL string) []Link {

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
                links = append(links, Link{URL: "https://en.wikipedia.org" + link})
            }
        }
    })

    c.Visit(pageURL)

    linkCache.Store(pageURL, links)

    return links
}

// normal
// func IDS(start, target string) []string {
//     depth := 0
//     for {
//         visited := make(map[string]bool)
//         path := make([]string, 0)
//         path = append(path, start)
//         path, found := DLS(path, visited, target, depth)
//         if found {
//             return path
//         }
//         depth++
//     }
// }

// pointer
// func IDS(start, target string) []string {
//     depth := 0
//     for {
//         visited := make(map[string]bool)
//         path := make([]string, 0)
//         path = append(path, start)
//         path, found := DLS(&path, &visited, target, depth)
//         if found {
//             return path
//         }
//         depth++
//     }
// }

// // sync map
func IDS(start, target string) ([]string, time.Duration) {
    startTime := time.Now()

    depth := 0
    for {
        visited := &sync.Map{}
        path := make([]string, 0)
        path = append(path, start)
        path, found := DLS(&path, visited, target, depth)
        if found {
            return path, time.Since(startTime)
        }
        depth++
    }
}

// normal
// func DLS(path []string, visited map[string]bool, target string, depth int) ([]string, bool) {
//     node := path[len(path)-1]

//     fmt.Println(path)

//     if node == target {
//         return path, true
//     }

//     if depth <= 0 {
//         return path, false
//     }

//     visited[node] = true

//     for _, link := range getLinks(node) {
//         if !visited[link.URL] {
//             newPath := append(path, link.URL)
//             newPath, found := DLS(newPath, visited, target, depth-1)
//             if found {
//                 return newPath, true
//             }
//         }
//     }

//     return path, false
// }

// pointer
// func DLS(path *[]string, visited *map[string]bool, target string, depth int) ([]string, bool) {
//     node := (*path)[len(*path)-1]

//     fmt.Println(*path)

//     if node == target {
//         return *path, true
//     }

//     if depth <= 0 {
//         return *path, false
//     }

//     (*visited)[node] = true

//     for _, link := range getLinks(node) {
//         if !(*visited)[link.URL] {
//             *path = append(*path, link.URL)
//             newPath, found := DLS(path, visited, target, depth-1)
//             if found {
//                 return newPath, true
//             }
//             // Undo the changes to the path and visited map after the recursive call
//             *path = (*path)[:len(*path)-1]
//             delete(*visited, link.URL)
//         }
//     }

//     return *path, false
// }

// sync map
func DLS(path *[]string, visited *sync.Map, target string, depth int) ([]string, bool) {
    node := (*path)[len(*path)-1]

    fmt.Println(*path)

    if node == target {
        return *path, true
    }

    if depth <= 0 {
        return *path, false
    }

    visited.Store(node, true)

    for _, link := range getLinks(node) {
        if _, ok := visited.Load(link.URL); !ok {
            *path = append(*path, link.URL)
            newPath, found := DLS(path, visited, target, depth-1)
            if found {
                return newPath, true
            }
            // Undo the changes to the path and visited map after the recursive call
            *path = (*path)[:len(*path)-1]
            visited.Delete(link.URL)
        }
    }

    return *path, false
}

func handler(w http.ResponseWriter, r *http.Request) {
    start := r.URL.Query().Get("start")
    target := r.URL.Query().Get("target")

    if start == "" || target == "" {
        http.Error(w, "Missing start or target parameter", http.StatusBadRequest)
        return
    }

    path, duration := IDS(start, target)

    if path != nil {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "path": path,
            "duration": duration,
        })
    } else {
        json.NewEncoder(w).Encode(map[string]interface{}{
            "error": "No path found",
        })
    }
}

func main() {
    // start := "https://en.wikipedia.org/wiki/Joko_Widodo"
    // target := "https://en.wikipedia.org/wiki/Sukarno"

    // path, duration := IDS(start, target)
    // fmt.Println("Execution time:", duration)


    // if path != nil {
    //     fmt.Println("Path: ", path)
    // } else {
    //     fmt.Println("No path found")
    // }

    mux := http.NewServeMux()
    mux.HandleFunc("/ids", handler)

    // Setup CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    })

    handler := c.Handler(mux)

    http.ListenAndServe(":8080", handler)
}