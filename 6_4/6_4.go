package main

import (
	"fmt"
)

type intSet struct {
	words []uint64
}

func (s *intSet) add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *intSet) addAll(x ...int) {
	for _, y := range x {
		s.add(y)
	}
}

func (s *intSet) elems() []int {
	results := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				results = append(results, 64*i+j)
			}
		}
	}
	return results
}

func main() {
	var x intSet
	x.addAll(357, 1, 32, 2, 64)
	fmt.Println("elems method in use: ", x.elems())
}
