package main

import (
	"fmt"
	"os"
	"strings"
)

// As of Go 1.16 "io/ioutil" is deprecated in favour of "os" and "io" pkgs

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
