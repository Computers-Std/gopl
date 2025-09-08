package main

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	// Echo1()
	// Echo2()
	Echo3()
}
