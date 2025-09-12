package shapes

import "math"

func Eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}
