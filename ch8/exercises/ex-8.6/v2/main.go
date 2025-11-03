package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"ukiran/gopl/ch5/links"
)

// Limit concurrent requests to 20
var tokens = make(chan struct{}, 20)

// crawl fetches links from a given URL
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	return list
}

// urlDepth keeps track of a URL and its crawl depth
type urlDepth struct {
	url   string
	depth int
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: crawler <url>")
	}

	maxDepth := 3
	worklist := make(chan urlDepth) // URLs to process
	var wg sync.WaitGroup

	seen := make(map[string]bool)
	var mu sync.Mutex

	// Start with command-line arguments
	for _, arg := range os.Args[1:] {
		wg.Add(1)
		go func(link string) {
			worklist <- urlDepth{link, 0}
		}(arg)
	}

	// Close worklist when done
	go func() {
		wg.Wait()
		close(worklist)
	}()

	// continuously receives values from the worklist channel, one by
	// one, until the channel is closed.
	for item := range worklist {
		// Skip if already seen
		mu.Lock()
		if seen[item.url] || item.depth >= maxDepth {
			mu.Unlock()
			wg.Done()
			continue
		}
		seen[item.url] = true
		mu.Unlock()

		// Crawl concurrently
		wg.Add(1)
		go func(u urlDepth) {
			defer wg.Done()
			for _, link := range crawl(u.url) {
				wg.Add(1)
				go func(link string) {
					worklist <- urlDepth{link, u.depth + 1}
				}(link)
			}
		}(item)
	}
}
