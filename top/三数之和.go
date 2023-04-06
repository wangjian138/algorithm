package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/3sum/submissions/
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	//先进行排序
	sort.Sort(sort.IntSlice(nums))

	start := 0
	end := len(nums) - 1

	var res [][]int

	for start < end {
		left := nums[start]
		twoSumTargetNum := twoSumTargetNew(nums[start+1:], 0-nums[start], nums[start])
		if len(twoSumTargetNum) > 0 {
			res = append(res, twoSumTargetNum...)
		}

		for start < end && nums[start] == left {
			start++
		}
	}

	return res
}

func twoSumTargetNew(nums []int, target, val int) [][]int {
	numsLen := len(nums)
	if numsLen == 0 {
		return [][]int{}
	}
	var res [][]int

	start := 0
	end := numsLen - 1

	for start < end {
		left := nums[start]
		right := nums[end]

		sums := left + right

		if sums < target {
			start++
			continue
		} else if sums > target {
			end--
		} else {
			res = append(res, []int{left, right, val})
			for start < end && nums[start] == left {
				start++
			}

			for start < end && nums[end] == right {
				end--
			}

		}
	}

	return res
}
