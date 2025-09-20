package main

import (
	"golang.org/x/net/html"
)

func forEachNode2(n *html.Node, pre, post func(n *html.Node) bool) (*html.Node, bool) {
	if pre != nil {
		if pre(n) {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node, ok := forEachNode2(c, pre, post); ok {
			return node, true
		}
	}
	if post != nil {
		if post(n) {
			return n, true
		}
	}
	return nil, false
}

func ElementById(doc *html.Node, id string) *html.Node {
	var preFunc = func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return true
			}
		}
		return false
	}
	if node, found := forEachNode2(doc, preFunc, nil); found {
		return node
	}
	return nil
}
