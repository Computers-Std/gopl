package main

import (
	"fmt"
	"log"
	"ukiran/gopl/ch7/exercises/ex-7.13/eval"
)

func main() {
	// expr, err := eval.Parse("sin(x*y/10)/10")
	expr, err := eval.Parse("pow(2,sin(y))*pow(2,sin(x))/12")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%v", expr)
}
