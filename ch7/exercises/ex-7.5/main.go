package main

import (
	"io"
	"log"
	"os"
)

// MyReader implements the io.Reader interface for a string
type MyReader struct {
	str string
	pos int
}

// NewMyReader creates a new MyReader for the given string.
func NewMyReader(s string) *MyReader {
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

type MyLimitedReader struct {
	R *MyReader // io.Reader
	N int
}

// Read reads data from MyReader, respecting the limit.
func (l *MyLimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}
	// Limit the length of p to the remaining bytes to read.
	if len(p) > l.N {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= n
	return
}

// MyLimitReader creates a MyLimitedReader with a limit.
func MyLimitReader(r *MyReader, n int) *MyLimitedReader {
	return &MyLimitedReader{r, n}
}

func main() {
	r := NewMyReader("some io.Reader stream to be read\n")
	lr := MyLimitReader(r, 4)

	// lr is of type "MyLimiterReader" which have a "MyReader" which
	// is a instance of io.Reader
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
