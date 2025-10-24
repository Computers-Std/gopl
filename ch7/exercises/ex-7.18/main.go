package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (n *Element) String() string {
	buf := &bytes.Buffer{}
	visit(n, buf, 0) // write to buf
	return buf.String()
}

func visit(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case *Element:
		fmt.Fprintf(w, "%*s%s %s\n", depth*2, "", n.Type.Local, n.Attr)
		for _, child := range n.Children {
			visit(child, w, depth+1)
		}
	case CharData:
		fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = elem
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, elem)
			}
			stack = append(stack, elem) // Push
		case xml.CharData:
			parent := stack[len(stack)-1]
			parent.Children = append(parent.Children, CharData(tok))
		case xml.EndElement:
			stack = stack[:len(stack)-1] // Pop
		}
	}
	return root, nil
}

func main() {
	node, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(node)
}
