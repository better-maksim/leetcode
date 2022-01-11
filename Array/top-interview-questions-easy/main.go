package main

import "fmt"

func removeDuplicates(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if len(nums) > i+1 {
			if nums[i] == nums[i+1] {
				nums = append(nums[:i], nums[i+1:]...)
				i = i -1
			}
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}
