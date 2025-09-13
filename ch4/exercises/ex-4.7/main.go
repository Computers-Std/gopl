// Exercise 4.7: Modify reverse to reverse the characters of a []byte
// slice that represents a UTF-8-encoded string , in place. Can you do
// it without allocating new memory?
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(bytes []byte) {
	buf := make([]byte, 0, len(bytes))
	i := len(bytes)

	for i > 0 {
		_, s := utf8.DecodeLastRune(bytes[:i])
		buf = append(buf, bytes[i-s:i]...)
		i -= s
	}
	copy(bytes, buf)
}

func main() {
	b := []byte("Hello, 世界")
	reverse(b)
	fmt.Println(string(b))
}

// NOTE: Need a revisit
