package main

import "fmt"

func main() {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	sub := minSubArrayLen(target, nums)
	fmt.Printf("minSubArrayLen:%v\n", sub)
}

// https://leetcode.cn/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	numsLen := len(nums)
	left, right := 0, numsLen-1

	index := 0
	sum := 0

	subLen := numsLen + 1

	for left <= right {
		sum += nums[left]
		left++
		for sum >= target {
			if left-index <= subLen {
				subLen = left - index
			}
			sum -= nums[index]
			index++
		}
	}
	if subLen == numsLen+1 {
		return 0
	}

	return subLen
}
