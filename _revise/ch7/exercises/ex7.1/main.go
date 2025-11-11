package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		(*c)++
	}
	return len(p), scanner.Err()
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		(*c)++
	}
	return len(p), scanner.Err()
}

func main() {
	var wc WordCounter
	wc.Write([]byte("hello world"))
	fmt.Println(wc)
	wc = 0
	var name = "Dolly"
	fmt.Fprintf(&wc, "hello, %s", name)
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("hello world\n good to see you"))
	fmt.Println(lc)
	lc = 0
	var line = "there is a way"
	fmt.Fprintf(&lc, "where there is a will\n%s", line)
	fmt.Println(lc)
}
