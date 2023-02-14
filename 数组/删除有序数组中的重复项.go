package main

import "fmt"

func main() {
	removeDuplicates([]int{
		0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
	})
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	slow := 0
	fast := 0

	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fmt.Printf("removeDuplicates 一样 slow:%v fast:%v\n", nums[slow], nums[fast])

			fast++
			continue
		}
		fmt.Printf("removeDuplicates slow:%v fast:%v\n", nums[slow], nums[fast])

		slow++
		nums[slow] = nums[fast]
		fast++

	}
	return slow + 1
}
