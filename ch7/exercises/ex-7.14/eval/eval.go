package eval

import (
	"fmt"
	"math"
)

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (m maxdivs) Eval(env Env) float64 {
	counts := make([]int, len(m.args))
	for i, a := range m.args {
		counts[i] = countDivs(int(a.Eval(env)))
	}

	max := 0
	for i, c := range counts {
		if c > counts[max] {
			max = i
		}
	}
	return m.args[max].Eval(env)
}

// Function to count divisors of a number
func countDivs(n int) int {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			count++
			if i != n/i { // Avoid double-counting the square root if n is a perfect square
				count++
			}
		}
	}
	return count
}
