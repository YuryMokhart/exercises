package main

import (
	"fmt"
	"sort"
)

type check []byte

func (c check) Len() int           { return len(c) }
func (c check) Less(i, j int) bool { return c[i] < c[j] }
func (c check) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func isPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("the string is palindrome:", isPalindrome(check([]byte("abba"))))
	fmt.Println("the string is palindrome:", isPalindrome(check([]byte("hello"))))
}
