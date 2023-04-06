package main

import (
	"fmt"
	"math"
)

func main() {
	minLen := lengthOfLongestSubstring("pwwkew")
	fmt.Printf("minLen:%v\n", minLen)
}

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	sMap := make(map[byte]int, len(s))

	slen := len(s)

	minLen := math.MinInt

	var (
		left  int
		right int
	)

	for right < slen {
		rByte := s[right]
		right++

		sMap[rByte]++

		if sMap[rByte] > 1 {
			lByte := s[left]
			sMap[lByte] = sMap[lByte] - 1
			left++
			minLen = getMax(right-left-1, minLen)

		}

	}
	return minLen
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
