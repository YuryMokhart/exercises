package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
	fmt.Println(phoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

//CreatePhoneNumber converts array into a sstring of the phone number format.
func CreatePhoneNumber(numbers [10]uint) string {
	strNumbers := fmt.Sprint(numbers)
	strNumbers = strings.ReplaceAll(strNumbers, " ", "")
	strNumbers = strings.Replace(strNumbers, "[", "", 1)
	strNumbers = strings.Replace(strNumbers, "]", "", 1)
	fmt.Println(strNumbers)
	phone := "(xxx) xxx-xxxx"
	for i := 0; i < len(strNumbers); i++ {
		phone = strings.Replace(phone, "x", string(strNumbers[i]), 1)
	}
	return phone
}
func phoneNumber(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}
