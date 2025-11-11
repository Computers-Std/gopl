package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
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
			builder.WriteString("_")
		}
		builder.WriteString(",")
		if t.right != nil {
			builder.WriteString(t.right.String())
		} else {
			builder.WriteString(")")
		}
		builder.WriteString(")")
	}
	return builder.String()
}

func main() {
	values := []int{4, 2, 5, 1, 3}

}
