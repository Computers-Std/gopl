package main

import (
	"unicode"
	"unicode/utf8"
)

func squashSpace(bytes []byte) []byte {
	out := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[1:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			/* ASCII space literal, untyped constant, will transfer to
			byte when assign to append's parameter ' ' is int32/rune,
			character in Go is actually Unicode code point, but string
			will utf8 encoded
			*/
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}

// NOTE: Not Original
