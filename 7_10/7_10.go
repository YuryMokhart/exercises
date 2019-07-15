package main

import (
	"fmt"
	"sort"
)

type check []byte

func (x check) Len() int           { return len(x) }
func (x check) Less(i, j int) bool { return x[i] < x[j] }
func (x check) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func isPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("the string is palindrome:", isPalindrome(check([]byte("abba"))))
	fmt.Println("the string is palindrome:", isPalindrome(check([]byte("hello"))))
}
