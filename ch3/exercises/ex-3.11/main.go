package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	var sign string
	if first := s[:1]; first == "+" || first == "-" {
		sign = first
		s = s[1:]
	}

	var prefixD, suffixD string
	if decimal := strings.LastIndex(s, "."); decimal != -1 {
		prefixD = s[:decimal]
		suffixD = s[decimal:]
	} else {
		prefixD = s
	}

	buf.WriteString(sign) // write Sign if any

	n := len(prefixD)
	if n > 3 {
		rest := n % 3
		if rest == 0 {
			rest = 3
		}
		buf.WriteString(prefixD[:rest])
		for i := rest; i < n; i += 3 {
			buf.WriteByte(',')
			buf.WriteString(prefixD[i : i+3])
		}
	} else {
		buf.WriteString(prefixD)
	}

	buf.WriteString(suffixD[:])
	return buf.String()

}

func main() {
	num1 := "-1234567"
	num2 := "+12345.6789"
	num3 := "145.6"
	fmt.Println(comma(num1))
	fmt.Println(comma(num2))
	fmt.Println(comma(num3))

}
