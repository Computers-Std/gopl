// Exercise 5.7: Develop startElement and endElement into a general
// HTML pretty-printer.  Print comment nodes, text nodes, and the
// attributes of each element (<a href='...'>). Use short forms like
// <img/> instead of <img></img> when an element has no
// children. Write a test to ensure that the output can be parsed
// successfully. (See Chapter 11.)
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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode, html.TextNode:
		if !(n.Parent.Type == html.ElementNode &&
			(n.Parent.Data == "script" || n.Parent.Data == "style")) {
			for line := range strings.SplitSeq(n.Data, "\n") {
				line = strings.TrimSpace(line)
				if line != "" && line != "\n" {
					fmt.Printf("%*s%s\n", depth*2, "", line)
				}
			}
		}
	case html.ElementNode:
		var attrs string
		for _, attr := range n.Attr {
			attrs += fmt.Sprintf("%s=%q ", attr.Key, attr.Val)
		}
		child := ""
		if n.Data == "img" && n.FirstChild == nil {
			child = " /"
		}
		if len(attrs) > 1 {
			attrs = attrs[:len(attrs)-1]
			fmt.Printf("%*s<%s %s%s>\n", depth*2, "", n.Data, attrs, child)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, child)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if !(n.Data == "img" && n.FirstChild == nil) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
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
