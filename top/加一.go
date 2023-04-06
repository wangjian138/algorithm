package main

import "fmt"

func main() {
	digits := []int{9, 9, 9}
	res := plusOne(digits)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/plus-one/
func plusOne(digits []int) []int {
	digitsLen := len(digits)

	for i := digitsLen - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digitsNew := make([]int, 0)
	digitsNew = append(digitsNew, 1)
	digitsNew = append(digitsNew, digits...)

	return digitsNew
}
