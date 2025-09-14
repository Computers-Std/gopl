package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Letter = iota
	Mark
	Digit
	Space
	Symbol
	UTFCategories
)

func main() {
	counts := make(map[rune]int) // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int
	var utfcat [UTFCategories]int
	invalid := 0

	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		switch {
		case unicode.IsLetter(r):
			utfcat[Letter]++
		case unicode.IsMark(r):
			utfcat[Mark]++
		case unicode.IsDigit(r):
			utfcat[Digit]++
		case unicode.IsSpace(r):
			utfcat[Space]++
		case unicode.IsSymbol(r):
			utfcat[Symbol]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nCategory   Count\n")
	// %m.gs, m: width, less<m print spacem >m,then show original char
	// g: control the count of char can be printed
	// -: left align, default: right align
	fmt.Printf("%-7.7s: %4d\n", "Letters", utfcat[Letter])
	fmt.Printf("%-7.7s: %4d\n", "Marks", utfcat[Mark])
	fmt.Printf("%-7.7s: %4d\n", "Digits", utfcat[Digit])
	fmt.Printf("%-7.7s: %4d\n", "Spaces", utfcat[Space])
	fmt.Printf("%-7.7s: %4d\n", "Symbols", utfcat[Symbol])

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
