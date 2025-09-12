package main

import "fmt"

func main() {
	res := isAnagram("listens", "silent")
	fmt.Println(res)
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount := make(map[rune]int)

	// count occurences in s1
	for _, char := range s1 {
		charCount[char]++
	}

	// subtract the occurences from s2
	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}
