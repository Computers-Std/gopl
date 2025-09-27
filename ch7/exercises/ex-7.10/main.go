package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for j > i {
		// Less() only
		// when i and j are equal
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	ints := []int{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(ints)))

	strings := []string{"racecar", "hello", "world", "hello", "racecar"}
	fmt.Println(IsPalindrome(sort.StringSlice(strings)))
}
