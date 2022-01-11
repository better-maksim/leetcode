package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

//时间复杂度 O(N)
func missingNumber(nums []int) int {
	xor :=0
	n := len(nums)
	for i := 0; i < n; i++ {
		xor ^= nums[i]
	}

	for i := 0; i <= n; i++ {
		xor ^= i
	}
	return xor
}
