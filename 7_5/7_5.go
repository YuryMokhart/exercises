package main

import (
	"fmt"
	"io"
	"strings"
)

type customLimitReader struct {
	r io.Reader
	n int64
}

func (r *customLimitReader) Read(result []byte) (n int, err error) {
	n, err = r.r.Read(result)
	if n > int(r.n) {
		n = int(r.n)
	}
	err = io.EOF
	return
}

func limitReader(r io.Reader, n int64) *customLimitReader {
	var clr customLimitReader
	clr.r = r
	clr.n = n
	return &clr
}

func main() {
	r := limitReader(strings.NewReader("That's a string, I need to read."), 16)
	result := make([]byte, 1024)
	n, err := r.Read(result)
	result = result[:n]
	fmt.Printf("Number of read bytes: %d\nerror: %s\nString: %s\n", n, err, string(result))
}
