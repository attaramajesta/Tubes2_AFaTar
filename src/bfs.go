package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    // Tentukan URL awal dan akhir
    startURL := "https://en.wikipedia.org/wiki/Toyota"
    targetURL := "https://en.wikipedia.org/wiki/Toyota_Corolla_Cross"

    // Panggil fungsi SolveWikiRace untuk menyelesaikan permainan Wiki Race
    links, err := SolveWikiRace(startURL, targetURL)
    if err != nil {
        log.Fatal(err)
    }

    // Tampilkan hasil
    fmt.Println("Links needed to go from", startURL, "to", targetURL+":")
    for _, link := range links {
        fmt.Println(link)
    }
}

// SolveWikiRace mencari jalur terpendek dari startURL ke targetURL menggunakan BFS
func SolveWikiRace(startURL, targetURL string) ([]string, error) {
    visited := make(map[string]bool)
    queue := [][]string{{startURL}}

    for len(queue) > 0 {
        // Ambil rute yang akan dievaluasi selanjutnya
        path := queue[0]
        queue = queue[1:]

        // Ambil URL terakhir dari rute
        currentURL := path[len(path)-1]

        // Jika kita sudah mencapai target, kembalikan jalur
        if currentURL == targetURL {
            return path, nil
        }

        // Jika URL sudah dikunjungi, lanjutkan ke rute berikutnya
        if visited[currentURL] {
            continue
        }

        // Ambil semua tautan pada halaman saat ini
        links, err := getLinks(currentURL)
        if err != nil {
            return nil, err
        }

        // Tambahkan tautan baru ke antrian untuk evaluasi berikutnya
        for _, link := range links {
            newPath := append([]string(nil), path...)
			fmt.Println(link)
            newPath = append(newPath, link)
            queue = append(queue, newPath)
        }

        // Tandai halaman saat ini sebagai sudah dikunjungi
        visited[currentURL] = true
    }

    // Jika tidak ada jalur yang ditemukan
    return nil, fmt.Errorf("no path found from %s to %s", startURL, targetURL)
}

// getLinks mengambil semua tautan dari halaman web yang diberikan
func getLinks(url string) ([]string, error) {
    var links []string

    // Lakukan HTTP GET request untuk halaman web
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Gunakan goquery untuk mengurai HTML
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return nil, err
    }

    // Temukan semua tautan pada halaman
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        // Ambil nilai atribut "href" dari tautan
        link, exists := s.Attr("href")
        if exists {
            // Filter tautan internal Wikipedia yang valid
            if strings.HasPrefix(link, "/wiki/") {
                // Hapus hash dari URL (tautan dalam halaman yang sama)
                link = strings.Split(link, "#")[0]

                // Hapus tautan ke file (misalnya gambar)
                if !strings.Contains(link, ":") {
                    links = append(links, "https://en.wikipedia.org"+link)
                }
            }
        }
    })

    return links, nil
}
