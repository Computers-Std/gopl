// non-recursive version of comma, using bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Comma(s string) string {
	var buf bytes.Buffer
	decimal := strings.LastIndex(s, ".")
	var prefix, suffix string

	if decimal != -1 {
		prefix = s[:decimal]
		suffix = "." + s[decimal+1:]
	} else {
		prefix = s
	}

	rest := len(prefix) % 3
	if rest == 0 {
		rest = 3
	}

	buf.WriteString(prefix[:rest])
	for i := rest; i < len(prefix); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(prefix[i : i+3])
	}
	buf.WriteString(suffix[:])
	return buf.String()

}

func main() {
	num1 := "1234567"
	num2 := "12345.6789"
	fmt.Println(Comma(num1))
	fmt.Println(Comma(num2))

}
