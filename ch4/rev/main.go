package main

import "fmt"

// reverse reverses a slice of ints in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a[:2]) // -> [1 0 2 3 4 5]
	reverse(a[2:]) // -> [1 0 5 4 3 2]
	reverse(a)     // -> [2 3 4 5 0 1]

	fmt.Println(a)
}
