import "math"

func polyColor(h int) srting {
	//
}

func polyHeight(x1, y1, x2, y2, x3, y3, x4, y4 float64) float64 {
	// ax + by + c => mx - y - (m*x1 - y1) = 0
	var linePointDist = func(a, b, c, x, y float64) float64 {
		numerator := math.Abs(a*x + b*y + c)
		denominator := math.Sqrt(a*a + b*b)
		return numerator / denominator
	}
	var m, a, b, c float64
	m = (y2 - y1) / (x2 - x1)
	a = m
	b = -1
	c = y1 - (m * x1)
	d1 := linePointDist(a, b, c, x3, y3)
	d2 := linePointDist(a, b, c, x4, y4)
	return math.Max(d1, d2)
}
