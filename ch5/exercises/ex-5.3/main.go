package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}

	scrapeText(doc)

}

func scrapeText(n *html.Node) {
	if n.Type == html.ElementNode &&
		(n.Data == "style" || n.Data == "script") {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		scrapeText(c)
	}
}
