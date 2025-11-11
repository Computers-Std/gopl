package main

import (
	"fmt"
	"strings"
)

// Tree struct definition
type tree struct {
	value       int
	left, right *tree
}

// Add inserts a new value into the tree
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

// String method to return a string representation of the tree
func (t *tree) String() string {
	if t == nil {
		return ""
	}

	// Use pre-order traversal to represent the tree as a string
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%d", t.value))

	// If there are subtrees, add them with parentheses
	if t.left != nil || t.right != nil {
		builder.WriteString("(")
		if t.left != nil {
			builder.WriteString(t.left.String()) // Recursively call String() on the left subtree
		} else {
			builder.WriteString("_") // If no left child, print "nil"
		}

		builder.WriteString(",")
		if t.right != nil {
			builder.WriteString(t.right.String()) // Recursively call String() on the right subtree
		} else {
			builder.WriteString("nil") // If no right child, print "nil"
		}
		builder.WriteString(")")
	}
	return builder.String()
}

// Sort function to build the binary search tree
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return root
}

func main() {
	// Example unsorted values
	values := []int{-1, 4, 2, 5, 1, 3}

	// Build the binary search tree
	root := Sort(values)
	fmt.Println(values)

	// Print the tree using the String() method
	fmt.Println(root) // This will implicitly call root.String()
}
