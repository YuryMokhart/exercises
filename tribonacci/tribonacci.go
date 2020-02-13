package main

import "fmt"

func main() {
	var slice [3]float64
	slice = [3]float64{300, 200, 100}
	s := Tribonacci(slice, 0)
	fmt.Println(s)
}

//Tribonacci represents a sequence of numbers where every number is equal to sum of 3 previous numbers.
func Tribonacci(signature [3]float64, n int) []float64 {
	var s []float64
	if n <= 0 {
		return s
	}
	if n < len(signature) && n > 0 {
		for i := 0; i < n; i++ {
			s = append(s, signature[i])
		}
		return s
	}
	if n >= len(signature) {
		for _, v := range signature {
			s = append(s, v)
		}
		for i := 3; i < n; i++ {
			s = append(s, s[i-1]+s[i-2]+s[i-3])
		}
	}
	return s
}
