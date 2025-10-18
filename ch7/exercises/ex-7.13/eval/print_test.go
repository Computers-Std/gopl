package eval

import (
	"fmt"
	"testing"
)

func TestSyntaxTree(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"sin(x*y/10)/10", "(sin(((x * y) / 10)) / 10)"},
		{"(sin(((x * y) / 10)) / 10)", "(sin(((x * y) / 10)) / 10)"},
		{"pow(2,sin(y))*pow(2,sin(x))/12", "((pow(2, sin(y)) * pow(2, sin(x))) / 12)"},
		{"((pow(2, sin(y)) * pow(2, sin(x))) / 12)", "((pow(2, sin(y)) * pow(2, sin(x))) / 12)"},
	}

	for _, test := range tests {
		expr, _ := Parse(test.input)
		got := fmt.Sprintf("%v", expr) // using .String()
		if got != test.want {
			t.Errorf("%s => %s, want %s",
				test.input, got, test.want)
		}
	}
}
