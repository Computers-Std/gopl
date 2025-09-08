// Exercise 2.1: Add types, constants, and functions to tempconv for
// processing temperatures in the Kelvin scale, where zero Kelvin is
// −273.15°C and a difference of 1K has the same magnitude as 1°C.
package main

import (
	"fmt"
	"ukiran/gopl/ch2/tempconv"
)

func main() {
	c := tempconv.FToC(212.0)
	fmt.Println(c.String()) // 100°C

	k := tempconv.KToC(5)
	fmt.Println(k.String()) // -268.5°C

	f := tempconv.CToF(11)
	fmt.Println(f.String()) // 51.8°F
}
