// Exercise 5.19: Use panic and recover to write a function that
// contains no return statement yet returns a non-zero value.
package main

import "fmt"

func recoverFromPanic() int {
	defer func() (result int) {
		// Recover from panic and set a value to return
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 42 // Set a non-zero return value
		}
		return
	}()

	// This could panic, but it doesn't need to return anything explicitly
	panic("something went wrong")
}

func main() {
	result := recoverFromPanic()
	fmt.Println("Recovered value:", result) // Output: Recovered value: 42
}

// [21-09-2025] NOTE: used ChatGPT
