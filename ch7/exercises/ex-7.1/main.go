package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		*c++
	}
	return len(p), s.Err()
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewReader(p))
	for s.Scan() {
		*c++
	}
	return len(p), s.Err()
}

func main() {
	var c WordCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 1 word
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 2 words

	var d LineCounter
	d.Write([]byte("hello 1\n hello 2\n hello again3\n"))
	fmt.Println(d) // 3 lines
	d = 0
	var lines = "Dolly\nPinky\nNimbu"
	fmt.Fprintf(&d, "how are you all!\n, %s", lines)
	fmt.Println(d) // 4 lines
}
