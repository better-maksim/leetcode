package main

import "fmt"

func main() {
	fmt.Println(romanToInt("DM"))
}

func romanToInt(s string) (ans int) {
	roman := map[byte]int{}

	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	fmt.Println(roman)

	n := len(s)

	for i := range s {

		value := roman[s[i]]
		if i < n-1 && value < roman[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return

}
