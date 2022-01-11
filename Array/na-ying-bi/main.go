package main

func minCount(coins []int) int {

	count := 0
	for _, value := range coins {
		if value%2 == 0 {
			count = count + value/2
		} else {
			count = count + value/2 + 1
		}
	}
	return count
}

func main() {
	minCount([]int{4, 2, 1})
}
