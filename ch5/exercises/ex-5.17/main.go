package main

import (
	"golang.org/x/net/html"
)

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

func ElementByTabName(doc *html.Node, names ...string) (result []*html.Node) {
	var preFunc = func(n *html.Node) {
		for _, tag := range names {
			if n.Type == html.ElementNode && n.Data == tag && n.FirstChild != nil {
				result = append(result, n.FirstChild)
			}
		}
	}
	forEachNode(doc, preFunc, nil)
	return
}
