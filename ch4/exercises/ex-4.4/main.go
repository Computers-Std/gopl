package main

import "fmt"

func rotate(s []int, n int) {
	var rot1 = func(s []int) {
		for i, j := 0, 1; j < len(s); i, j = i+1, j+1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	for i := 0; i < n; i++ {
		rot1(s)
	}
}

func rot1v2(s []int) {
	first := s[0]
	copy(s, s[1:])
	s[len(s)-1] = first
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println(a)
}
