package main

import (
	"fmt"
	"os"
)

func echoEx1() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func echoEx2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%v -> %v\n", i+1, arg)
	}
}

func main() {
	echoEx1()
	echoEx2()
}
