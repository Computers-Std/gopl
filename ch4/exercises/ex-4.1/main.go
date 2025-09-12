package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n%T\n\n", c1, c2, c1 == c2, c1, c2)

	fmt.Println(bitsDiff(&c1, &c2))
}

// difference means 'XOR'
// only accepted sha256 digest
func bitsDiff(b1, b2 *[sha256.Size]byte) int {
	var sum int
	for i := 0; i < sha256.Size; i++ {
		sum += int(pc[b1[i]^b2[i]])
	}
	return sum
}
