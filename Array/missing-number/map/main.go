package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

//时间复杂度 O(N)
func missingNumber(nums []int) int {
	arrLen := len(nums)
	arr := make(map[int]int)

	for i, v := range nums {
		arr[v] = i
	}

	for i := 1; i <= arrLen; i++ {
		_, ok := arr[i]
		if !ok {
			return i
		}
	}
	return 0
}
