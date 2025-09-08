// Exercise 1.4: Modify dup2 to print the names of all files in which
// each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		fileName := "Stdin"
		countLines(os.Stdin, counts, fileName)
	} else {
		for _, fn := range files {
			f, err := os.Open(fn)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fn)
			f.Close()
		}
	}

	for line, repititions := range counts {
		n := len(repititions)
		allFiles := slices.Compact(repititions) // Uniq values
		if n > 1 {
			fmt.Printf("%d %v\t%s\n", n, allFiles, line)
		}
	}
}

func countLines(f *os.File, counts map[string][]string, fn string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text] = append(counts[text], fn)
	}
}
