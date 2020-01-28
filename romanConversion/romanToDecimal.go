package main

import "fmt"

func main() {
	fmt.Println(number("MCDXV"))
	Decode("MCDXV")

}

func symbolConvert(symbol string) int {
	if symbol == "I" {
		return 1
	}
	if symbol == "V" {
		return 5
	}
	if symbol == "X" {
		return 10
	}
	if symbol == "L" {
		return 50
	}
	if symbol == "C" {
		return 100
	}
	if symbol == "D" {
		return 500
	}
	if symbol == "M" {
		return 1000
	}
	return -1
}

func number(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		s1 := symbolConvert(string(str[i]))
		if i+1 < len(str) {
			s2 := symbolConvert(string(str[i+1]))
			if s1 >= s2 {
				result = result + s1
			} else {
				result = result + s2 - s1
				i++
			}
		} else {
			result = result + s1
		}
	}
	return result
}

var mapping = map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

//Decode converts Roman numerals to decimals.
func Decode(roman string) int {
	var intern []int
	var result int
	for _, r := range roman {
		intern = append(intern, mapping[r])
		fmt.Println(intern)
	}

	for i := 1; i < len(intern); i++ {
		if intern[i-1] >= intern[i] {
			result += intern[i-1]
			fmt.Println(result)
		} else {
			result -= intern[i-1]
			fmt.Println(result)
		}
	}

	result += intern[len(intern)-1]
	fmt.Println(result)

	return result
}
