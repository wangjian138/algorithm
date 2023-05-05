package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	res := summaryRanges(nums)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/summary-ranges/?envType=study-plan-v2&id=top-interview-150
func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	left := 0
	right := len(nums)

	for left < right {
		num := nums[left]
		temp := num

		for left+1 < right && nums[left+1] == num+1 {
			left++
			num = nums[left]
		}
		if temp == num {
			res = append(res, strconv.Itoa(num))
			left++
			continue
		}
		res = append(res, fmt.Sprintf("%v->%v", temp, nums[left]))
		left++
	}

	return res
}
