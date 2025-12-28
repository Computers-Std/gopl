// Exercise 8.10: HTTP requests may be cancelled by closing the
// optional Cancel channel in the http.Request struct. Modify the
// webcrawler of Section 8.6 to support cancellation.

// Hint: the http.Get convenience function does not give you an
// opportunity to customize a Request. Instead, create the request
// using http.NewRequest, set its Cancel field, then perform the
// request by calling http.DefaultClient.Do(req).
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

var ctx, cancel = context.WithCancel(context.Background())

func cancelled() bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// ch5/links
// Extract makes an HTTP GET request to the specified URL, parses the
// response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad urls
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

// ch8/crawl
func crawl(url string) []string {
	if cancelled() {
		return nil
	}
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		// If the error was due to cancellation, we don't need to log it.
		if ctx.Err() == nil {
			log.Print(err)
		}
	}
	return list
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel()                       // This closes the ctx.Done() channel
	}()

	var wg sync.WaitGroup
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// Add command-line arguments to worklist
	initialLinks := os.Args[1:]
	if len(initialLinks) > 0 {
		wg.Add(len(initialLinks))
		go func() { worklist <- initialLinks }()
	}

	// Monitor for completion
	go func() {
		wg.Wait()
		cancel()
	}()

	// Create 20 crawler goroutines to fetch each unseen link
	for range 20 {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					select {
					case worklist <- foundLinks:
					case <-ctx.Done():
						return
					}
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items and sends the
	// unseen ones to the crawlers
	seen := make(map[string]bool)
	for {
		select {
		case <-ctx.Done():
			return
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					wg.Go(func() {
						unseenLinks <- link
					})
				}
			}
		}
	}
}
