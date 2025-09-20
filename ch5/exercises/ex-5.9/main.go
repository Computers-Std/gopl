package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	var result string
	for word := range strings.SplitSeq(s, " ") {
		after, found := strings.CutPrefix(word, "$")
		if found {
			result += f(after)
		} else {
			result += word
		}
		result += " "
	}
	return strings.TrimSpace(result)
}

func main() {
	f := func(s string) string {
		return "[" + s + "]"
	}
	expanded := expand("Hello $world this is $Go", f)
	fmt.Println(expanded)
}
