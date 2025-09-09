package convert

import (
	"fmt"
	"math"
)

type (
	Celsius    float64
	Fahrenheit float64
	Meters     float64
	Feet       float64
	Pounds     float64
	Kilograms  float64
)

const (
	oneMeterToFeet     = 3.28084
	oneFeetToMeter     = 0.3048
	onePoundToKilogram = 0.453592
	oneKilogramToPound = 2.20462
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%g m", m)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%g ft", ft)
}

func (lb Pounds) String() string {
	return fmt.Sprintf("%g lbs", lb)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%g kg", k)
}

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// MToF converts a Meters scale length to Feet
func MToF(m Meters) Feet {
	return Feet(math.Abs(float64(m)) * oneMeterToFeet)
}

// FToM converts a Feet scale length to Meters
func FToM(ft Feet) Meters {
	return Meters(math.Abs(float64(ft)) * oneFeetToMeter)
}

// PToK converts a Pounds scale weight to Kilograms
func PToK(lb Pounds) Kilograms {
	return Kilograms(math.Abs(float64(lb)) * onePoundToKilogram)
}

// KToP converts a Kilograms scale weight to Pounds
func KToP(k Kilograms) Pounds {
	return Pounds(math.Abs(float64(k)) * oneKilogramToPound)
}
