// The strings.NewReader function returns a value that satisfies the
// io.Reader interface (and others) by reading from its argument, a
// string . Implement a simple version of NewReader yourself,

package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// MyReader implements the io.Reader interface for a string
type MyReader struct {
	str string
	pos int
}

// MyNewReader creates a new MyReader for the given string.
func MyNewReader(s string) *MyReader {
	return &MyReader{s, 0}
}

// Read reads data from the string into the given buffer.
func (r *MyReader) Read(b []byte) (n int, err error) {
	if r.pos >= len(r.str) {
		return 0, io.EOF
	}
	n = copy(b, r.str[r.pos:])
	r.pos += n
	return
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(MyNewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
	}
}
