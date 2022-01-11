package main

func twoSum(nums []int, target int) []int {
	numLen := len(nums)
	for i := 0; i < numLen; i++ {
		for j := i + 1; j < numLen; j++ {
			if nums[i] == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func hashTowSUm(nums []int, target int) []int {
	hasTable := map[int]int{}
	for idx, value := range nums {
		if p, ok := hasTable[target-value]; ok {
			return []int{p, idx}
		}
		hasTable[value] = idx
	}
	return nil
}

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)

}
