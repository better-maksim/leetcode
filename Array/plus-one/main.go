package main

import "fmt"

func plusOne(digits []int) []int {


	l := len(digits)


	for i := l - 1; i >= 0; i-- {

		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			return digits
		}
	}

	temp := make([]int,l + 1)
	temp[0] = 1
	fmt.Println(temp)
	return temp
}

func main() {
	plusOne([]int{0})
}
