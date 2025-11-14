package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%v)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%v %c %v)", b.x, b.op, b.y)
}

func (c call) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		write(buf, arg)
	}
	buf.WriteByte(')')
	return buf.String()
}
