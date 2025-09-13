package main

import "fmt"

func rmDuplicates(strs []string) []string {
	out := strs[:0]
	for i := 0; i < len(strs); i++ {
		if i+1 == len(strs) || i+1 < len(strs) && strs[i] != strs[i+1] {
			out = append(out, strs[i])
		}
	}
	return out
}

func rmDuplicates2(str []string) []string {
	out := str[:0]
	for i := 0; i < len(str); i++ {
		if i == 0 || str[i] != str[i-1] {
			out = append(out, str[i])
		}
	}
	return out
}

func rmDuplicates3(strs []string) []string {
	w := 0
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}

func main() {
	s1 := []string{"h", "e", "e", "l", "l", "o", "o"}
	s2 := []string{"h", "e", "e", "l", "l", "o", "o"}
	s3 := []string{"h", "e", "e", "l", "l", "o", "o"}
	fmt.Println(rmDuplicates(s1))
	fmt.Println(rmDuplicates2(s2))
	fmt.Println(rmDuplicates3(s3))
}
