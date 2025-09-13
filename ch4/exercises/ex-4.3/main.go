package main

import "fmt"

// reverse an array
func reverse(a *[6]int) {
	n := len(a) - 1
	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func reverse2(a *[6]int) {
	for i := 0; i < len(a)/2; i++ {
		tail := len(a) - i - 1
		a[i], a[tail] = a[tail], a[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5} // Initialize the array
	reverse(&a)                     // Pass a pointer to reverse
	fmt.Println(a)                  // Print the reversed array
}
