package main

import "io"

type sReader struct {
	str string
	pos int
}

func (r *sReader) Read(b []byte) (n int, err error) {
	if r.pos >= len(r.str) {
		return 0, io.EOF
	}
	n = copy(b, r.str[r.pos:])
	r.pos += n
	return
}

func StrNewReader(s string) *sReader {
	return &sReader{s, 0}
}
