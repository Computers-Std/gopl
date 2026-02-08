package main

import (
	"fmt"
	"io"
	"os"
)

type CounterWriter struct {
	counter int64
	writer  io.Writer
}

// must be a pointer type in order to count
func (c *CounterWriter) Write(p []byte) (int, error) {
	c.counter += int64(len(p))
	return c.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{0, w}
	return &cw, &cw.counter
}

// CounterWriter is a proxy
func main() {
	cw, c := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "This is %dth day with a smile on face.\n", 26)
	fmt.Println(*c)

	// another output
	fmt.Fprint(cw, "ha ha...\n")
	fmt.Println(*c)
}
