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

func (s *intSet) unionWith(t *intSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
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

func (s *intSet) len() (length int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return
}

func (s *intSet) remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
}

func (s *intSet) clear() {
	s.words = s.words[:0]
}

func (s *intSet) copy() *intSet {
	var n intSet
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return &n
}

func main() {
	var x, y intSet
	x.add(1)
	x.add(144)
	x.add(9)
	fmt.Println("x string: ", x.string())
	y.add(9)
	y.add(42)
	fmt.Println("y string: ", y.string())
	fmt.Println("the length of the x string: ", x.len())
	x.remove(144)
	fmt.Println("x string after deleting 144: ", x.string())
	z := x.copy()
	fmt.Println("copied string is ", z.string())
	x.unionWith(&y)
	fmt.Println("united string: ", x.string())
	fmt.Printf("the set contains number 9: %t\nthe set contains number 123: %t\n", x.has(9), x.has(123))
	x.clear()
	fmt.Println("empty set", x.string())
}
