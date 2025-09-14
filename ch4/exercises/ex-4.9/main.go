// wordfreq reports frequency of each word in an input text file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	frequency := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		frequency[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	for w, f := range frequency {
		fmt.Println(w, f)
	}
}
