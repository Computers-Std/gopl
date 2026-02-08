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
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	cw := &countingWriter{writer: w, count: &count}
	return cw, &count
}

func main() {
	writer, count := CountingWriter(os.Stdout)

	writer.Write([]byte("Hello, world.\n"))
	writer.Write([]byte("Go is awesome!\n"))

	fmt.Println("Bytes written:", *count)
}
