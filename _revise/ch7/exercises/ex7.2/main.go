package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	n, err = cw.writer.Write(p) // write to original writer
	*cw.count += int64(n)       // update the count
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	cw := &countingWriter{writer: w, count: &count}
	return cw, &count
}

// example usage
func main() {
	f, err := os.Create("out.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	// CountingWriter wraps around the file writer
	writer, count := CountingWriter(f)

	writer.Write([]byte("Hello, world,"))
	writer.Write([]byte(" Go is awesome!"))

	fmt.Println("Bytes written:", *count)
}
