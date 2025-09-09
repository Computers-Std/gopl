package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Example number
	num := 3

	// Convert to binary
	binaryStr := strconv.FormatInt(int64(num), 2)

	// Pad with leading zeroes to 8 bits
	desiredLength := 8
	binaryStr = fmt.Sprintf("%0*s", desiredLength, binaryStr)

	// Output the result
	fmt.Println(binaryStr) // Output: 00000011
}
