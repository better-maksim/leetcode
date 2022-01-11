package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	arrSum := 0
	for _, num := range nums {
		arrSum += num
	}
	return total - arrSum
}
