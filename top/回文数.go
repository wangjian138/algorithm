package main

import "strconv"

func main() {

}

// https://leetcode.cn/problems/palindrome-number/?envType=study-plan-v2&id=top-interview-150
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	str := strconv.Itoa(x)

	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
