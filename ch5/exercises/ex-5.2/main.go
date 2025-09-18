package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	var elements = make(map[string]int)
	mapper(elements, doc)

	// out
	for n, elem := range elements {
		fmt.Println(n, elem)
	}

}

func mapper(amap map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		amap[n.Data]++
	}

	if n.FirstChild != nil {
		mapper(amap, n.FirstChild)
	}
	if n.NextSibling != nil {
		mapper(amap, n.NextSibling)
	}
}
