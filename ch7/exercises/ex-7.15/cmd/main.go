package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"ukiran/gopl/ch7/exercises/ex-7.15/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined varaible: %s", v)
		}
	}
	return expr, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the Expression: ")
	scanner.Scan()
	expression := scanner.Text()
	fmt.Print("Enter Args(CRM): ")
	scanner.Scan()
	args := strings.Split(scanner.Text(), ",")
	if len(args) > 3 {
		log.Fatalf("Given %v vars, but program is designed for atmost 3 vars", len(args))
	}

	expr, err := parseAndCheck(expression)
	if err != nil {
		log.Fatal(err)
	}

	argmap := make(map[string]int)
	vars := []string{"x", "y", "z"}
	for i, a := range args {
		val, err := strconv.Atoi(strings.Trim(a, " "))
		if err != nil {
			log.Fatalf("Invalid argument %v: %v", a, err)
		}
		argmap[vars[i]] = val
	}
	for i := len(args); i < len(vars); i++ {
		argmap[vars[i]] = 0
	}

	env := eval.Env{
		"x": float64(argmap["x"]),
		"y": float64(argmap["y"]),
		"z": float64(argmap["z"]),
	}

	fmt.Printf("\n%v = %v\n", expression, expr.Eval(env))
}
