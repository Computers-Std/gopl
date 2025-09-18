package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(countWordsAndImages(doc))
}

func countWordsAndImages(n *html.Node) (words, images int) {
	switch n.Type {
	case html.ElementNode:
		if n.Data == "style" || n.Data == "script" {
			return
		} else if n.Data == "img" {
			images++
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		for _, line := range strings.Split(text, "\n") {
			if line != "" {
				words += len(strings.Split(line, " "))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}
