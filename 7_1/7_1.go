package main

import (
	"bufio"
	"fmt"
	"strings"
)

type wordCounter int

func (wc *wordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, nil
}

type lineCounter int

func (lc *lineCounter) Write(p []byte) (int, error) {
	count := 1
	for _, b := range p {
		if b == '\n' {
			count++
		}
	}
	return count, nil
}

func main() {
	var wc wordCounter
	words, err := wc.Write([]byte("Hello, it's me!"))
	if err != nil {
		fmt.Printf("write words counter error occured %s", err)
	}
	fmt.Println(words)
	var lc lineCounter
	lines, err := lc.Write([]byte(`hello
	it
	is
	me`))
	if err != nil {
		fmt.Printf("write lines counter error occured %s", err)
	}
	fmt.Println(lines)
}
