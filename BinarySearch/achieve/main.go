package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(rank(5, nums))
}

func rank(key int, nums []int) int {
	//nums 必须是有序的，因为要通过游标的方式进行缩小范围，如果没有续

	lo := 0             //游标的地位
	hi := len(nums) - 1 //游标的高位

	for {
		//通过循环判断来不断的移动游标的位置

		if lo > hi {
			break
		}
		//计算中间值，因为是二分查找，首先就直接上来砍掉一半的数据
		mid := lo + (hi-lo)/2


		if key < nums[mid] { //如果中间值小于中阿金状态，那么久开始进行移动
			hi = mid - 1
		} else if key > nums[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
