package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
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

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%d", t.value))
	if t.left != nil || t.right != nil {
		builder.WriteString("(")
		if t.left != nil {
			builder.WriteString(t.left.String())
		} else {
			builder.WriteString("nil")
		}
		builder.WriteString(",")
		if t.right != nil {
			builder.WriteString(t.right.String())
		} else {
			builder.WriteString("nil")
		}
		builder.WriteString(")")
	}
	return builder.String()
}

func makeTree(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	return root
}

func main() {
	values := []int{4, 2, 5, 1, 3}
	root := makeTree(values)
	fmt.Println(root)
}
