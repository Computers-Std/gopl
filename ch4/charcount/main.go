// Charcount computes counts of Unicode characters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts, utflen, invalid, err := countData(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("rune\tcount\n")

	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func countData(r io.Reader) (map[rune]int, [utf8.UTFMax + 1]int, int, error) {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	reader := bufio.NewReader(r)
	for {
		// ReadRune returns the rune, its size in bytes, and any error
		runeVal, size, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, utflen, 0, err
		}

		// If ReadRune encounters an invalid UTF-8 sequence, it returns
		// unicode.ReplacementChar with a size of 1 byte.
		if runeVal == unicode.ReplacementChar && size == 1 {
			invalid++
			continue
		}

		counts[runeVal]++
		utflen[size]++
	}
	return counts, utflen, invalid, nil
}
