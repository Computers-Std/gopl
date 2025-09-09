// Exercise 2.2: Write a general-purpose unit-conversion program
// analogous to cf that reads numbers from its command-line arguments
// or from the standard input if there are no arguments, and converts
// each number into units like temperature in Celsiu s and Fahrenheit,
// length in feet and meters, weight in pound s and kilograms, and the
// like.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"ukiran/gopl/ch2/exercises/ex-2.2/convert"
)

func main() {
	var values []float64
	args := os.Args[1:]

	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			n, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "strconv: %v\n", err)
				os.Exit(1)
			}
			values = append(values, n)
		}
	} else {
		for _, arg := range os.Args[1:] {
			n, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "strconv: %v\n", err)
				os.Exit(1)
			}
			values = append(values, n)
		}
	}
	printResults(values)
}

func printResults(values []float64) {
	newline := "-------\n"
	for _, n := range values {
		f := convert.Fahrenheit(n)
		c := convert.Celsius(n)
		m := convert.Meters(math.Abs(n))
		ft := convert.Feet(math.Abs(n))
		lb := convert.Pounds(math.Abs(n))
		k := convert.Kilograms(math.Abs(n))
		fmt.Printf(" %s = %s, %s = %s\n %s = %s, %s = %s\n %s = %s, %s = %s\n %s",
			f, convert.FToC(f),
			c, convert.CToF(c),
			m, convert.MToF(m),
			ft, convert.FToM(ft),
			lb, convert.PToK(lb),
			k, convert.KToP(k),
			newline,
		)
	}
}
