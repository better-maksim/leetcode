package main

import "fmt"

func main() {
	e1 := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(e1 == 6)

}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
