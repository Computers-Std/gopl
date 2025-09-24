package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// appendValues appends the elements of t to values in order and
// returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// String method for *tree
// The String method is a special method in Go that is called when you
// try to print or format a value (e.g., fmt.Println(t)).
func (t *tree) String() string {
	if t == nil {
		return ""
	}
	values := inOrderTraversal(t)
	// join the values into a string with commas
	return fmt.Sprintf("[%s]", strings.Join(values, ","))
}

func inOrderTraversal(t *tree) []string {
	if t == nil {
		return nil
	}

	left := inOrderTraversal(t.left)
	right := inOrderTraversal(t.right)
	// convert the current node value to string
	current := fmt.Sprintf("%d", t.value)

	result := append(left, current)
	result = append(result, right...)
	return result
}

func main() {
	// Creating a simple tree
	var root *tree
	root = add(root, 5)
	root = add(root, 3)
	root = add(root, 8)
	root = add(root, 2)
	root = add(root, 4)
	root = add(root, 7)
	root = add(root, 9)

	// Printing the tree using the String method
	// fmt.Println(root.String()) // instead
	fmt.Println(root) // This will call the String method of the tree

}
