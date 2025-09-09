package main

import (
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of 1.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var result byte
	for i := 0; i < 8; i++ {
		result += pc[byte(x>>(i*8))]
		// result += pc[byte(x>>uint(i))]
	}

	return int(result)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number")
		return
	}

	arg := os.Args[1]
	num, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		fmt.Printf("Invalid number: %v\n", err)
		return
	}
	r := PopCount(num)
	fmt.Println("Population count: ", r)
}
