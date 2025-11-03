package main

import (
	"fmt"
	"log"
	"os"
	"ukiran/gopl/ch5/links"
)

// tokens is counting semaphore used to enforce a limit of 20
// concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // aquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n += 1
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	// Crawl the web concurrently
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n += 1
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
