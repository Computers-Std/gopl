package main

import (
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 12345.6789 => 12,345.6789
func commaF(s string) string {
	decimal := strings.LastIndex(s, ".")
	s0 := s[:decimal]
	// return s0
	return comma(s0) + "." + s[decimal+1:]
}
