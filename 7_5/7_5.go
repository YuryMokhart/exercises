package main

import (
	"fmt"
	"io"
	"strings"
)

type customLimitReader struct {
	r io.Reader
	n int
}

func (r *customLimitReader) Read(result []byte) (n int, err error) {
	n, err = r.r.Read(result)
	if n > r.n {
		n = r.n
	}
	err = io.EOF
	return n, err
}

func limitReader(r io.Reader, n int) *customLimitReader {
	var clr customLimitReader
	clr.r = r
	clr.n = n
	return &clr
}

func main() {
	cr := limitReader(strings.NewReader("That's a string, I need to read."), 16)
	result := make([]byte, 1024)
	n, err := cr.Read(result)
	result = result[:n]
	fmt.Printf("String: %s\nerror: %s\nNumber of read bytes: %d\n", string(result), err, n)
}
