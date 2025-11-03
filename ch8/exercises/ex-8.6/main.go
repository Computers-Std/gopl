package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"ukiran/gopl/ch5/links"
)

// Limit concurrent requests to 20
var tokens = make(chan struct{}, 20)

// crawl fetches links from a URL
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire token
	list, err := links.Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	return list
}

// urlDepth pairs a URL with its depth level.
type urlDepth struct {
	url   string
	depth int
}

func main() {
	maxDepth := flag.Int("depth", 2, "crawl depth")
	flag.Parse()

	worklist := make(chan []urlDepth)
	var n int // number of pending sends to worklist

	// Start with command-line arguments at depth 0
	n++
	go func() {
		start := os.Args[1:]
		var list []urlDepth
		for _, url := range start {
			list = append(list, urlDepth{url, 0})
		}
		worklist <- list
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		items := <-worklist
		for _, item := range items {
			if !seen[item.url] && item.depth < *maxDepth {
				seen[item.url] = true
				n++
				go func(u urlDepth) {
					links := crawl(u.url)
					var next []urlDepth
					for _, link := range links {
						next = append(next, urlDepth{link, u.depth + 1})
					}
					worklist <- next
				}(item)
			}
		}
	}
}
