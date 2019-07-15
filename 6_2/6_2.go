package main

import (
	"bytes"
	"fmt"
)

type intSet struct {
	words []uint64
}

func (s *intSet) has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *intSet) add(x int) {
	word, bit := x/64, uint(x%64)
	fmt.Println("word = ", word, "   bit = ", bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *intSet) string() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *intSet) addAll(x ...int) {
	for _, y := range x {
		s.add(y)
	}
}

func main() {
	var x intSet
	x.addAll(1, 14, 82)
	fmt.Println("addAll function in use: ", x.string())
}
