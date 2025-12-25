package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type urlDepth struct {
	url   string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	maxDepth := flag.Int("depth", 2, "crawl depth")
	flag.Parse()

	worklist := make(chan []urlDepth)
	unseenItems := make(chan urlDepth)

	// Start with the command-line arguments
	initialLinks := []urlDepth{}
	for _, url := range flag.Args() {
		initialLinks = append(initialLinks, urlDepth{url, 0})
	}

	// 1. Start workers first so they are ready to receive
	for i := 0; i < 20; i++ {
		go func() {
			for item := range unseenItems {
				foundLinks := crawl(item.url)
				var foundItems []urlDepth
				for _, link := range foundLinks {
					foundItems = append(foundItems, urlDepth{link, item.depth + 1})
				}
				// Send back to worklist in a goroutine to avoid blocking
				go func() { worklist <- foundItems }()
			}
		}()
	}

	// 2. The coordinator loop
	go func() { worklist <- initialLinks }()

	seen := make(map[string]bool)

	// Because the work is finite (depth-limited), we use a counter to
	// exit the loop when all pending tasks are finished
	n := 1 // Number of pending sends to worklist
	for ; n > 0; n-- {
		list := <-worklist
		for _, item := range list {
			if !seen[item.url] && item.depth <= *maxDepth {
				seen[item.url] = true
				n++ // increment for every new work item we're about to process
				unseenItems <- item
			}
		}
	}
}
