package main

import "fmt"

func main() {
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
}

// FindOutlier takes the array as an argument and returns this "outlier".
func FindOutlier(integers []int) int {
	var even []int
	var odd []int
	for _, v := range integers {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	if len(even) > len(odd) {
		return odd[0]
	}
	if len(even) < len(odd) {
		return even[0]
	}
	return 0
}
