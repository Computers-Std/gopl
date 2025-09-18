package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*
<a href="https://www.example.com">Visit Example</a>
<img src="image.jpg" alt="Description of Image">
<script src="script.js"></script>
<link rel="stylesheet" href="styles.css">
*/
var mapping = map[string]string{
	"a":      "href", // "http"
	"img":    "src",  // "."
	"script": "src",  // ".js"
	"link":   "href", // "." or ".css" but, "stylesheet" needs nesting
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Anchor links:")
	for _, link := range visit2("a", "http", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("\nImage links:")
	for _, link := range visit2("img", ".", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("\nScript links:")
	for _, link := range visit2("script", ".js", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("\nStylesheet links:")
	for _, link := range visit2("link", ".", nil, doc) {
		fmt.Println(link)
	}
}

func visit2(target string, match string, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == target {
		for _, a := range n.Attr {
			if a.Key == mapping[target] && strings.Contains(a.Val, match) {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit2(target, match, links, c)
	}
	return links
}
