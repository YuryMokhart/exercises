package main

import (
	"fmt"
	"math"
)

func main() {
	n := 76576500
	for {
		divQ := divide(n)
		if divQ >= 500 && isTriangle(n) == true {
			fmt.Printf("the number %d is a triangular number! woohoo!\n", n)
			break
		}
		n++
	}
}

func divide(number int) int {
	divQuantity := 0
	for i := 1; i*i <= number; i++ {
		if number%i == 0 {
			divQuantity += 2
		}
	}
	return divQuantity
}

func isTriangle(num int) bool {
	if num < 0 {
		return false
	}
	n := (math.Sqrt(1+8*float64(num)) - 1) / 2
	if n-float64(int(n)) == 0 {
		return true
	}
	return false
}
