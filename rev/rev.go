package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords("hello mate"))
}

//SpinWords takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed.
func SpinWords(str string) string {
	s := strings.Split(str, " ")
	for i, v := range s {
		if len(v) >= 5 {
			chars := []rune(v)
			for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
				chars[i], chars[j] = chars[j], chars[i]
			}
			s[i] = string(chars)
		}
	}
	str = fmt.Sprint(s)
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	return str
}
