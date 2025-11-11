package main

import "io"

type LimitedReader struct {
	reader io.Reader
	limit  int64
}

func (r *LimitedReader) Read(b []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(b)) > r.limit {
		b = b[0:r.limit]
	}

	n, err = r.reader.Read(b) // read n bytes

	// when len(b) is less than r.limit, then subtract n(read bytes)
	// from the r.limit
	r.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}
