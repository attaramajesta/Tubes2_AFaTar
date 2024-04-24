package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly"
    "strings"
)

// Link represents a link between Wikipedia pages
type Link struct {
    URL string
}

// getWikiLinks function fetches links from a Wikipedia page
func getWikiLinks(pageURL string) []Link {
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

    return links
}


func IDS(start, target string, maxDepth int) []string {
    for depth := 0; depth <= maxDepth; depth++ {
        visited := make(map[string]bool)
        path := make([]string, 0)
        path = append(path, start)
        path, found := DLS(path, visited, target, depth)
        if found {
            return path
        }
    }

    return nil
}

func DLS(path []string, visited map[string]bool, target string, depth int) ([]string, bool) {
    node := path[len(path)-1]

    // fmt.Println(path)

    if node == target {
        return path, true
    }

    if depth <= 0 {
        return path, false
    }

    visited[node] = true

    for _, link := range getWikiLinks(node) {
        if !visited[link.URL] {
            newPath := append(path, link.URL)
            newPath, found := DLS(newPath, visited, target, depth-1)
            if found {
                return newPath, true
            }
        }
    }

    return path, false
}

func main() {
    start := "https://en.wikipedia.org/wiki/Joko_Widodo"
    target := "https://en.wikipedia.org/wiki/Jusuf_Kalla"

    start_time := time.Now()
    path := IDS(start, target, 1)
    fmt.Println("Execution time:", time.Since(start_time))


    if path != nil {
        fmt.Println("Path: ", path)
    } else {
        fmt.Println("No path found")
    }
}