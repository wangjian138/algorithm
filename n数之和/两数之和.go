package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int][]int, len(nums))
	for k, v := range nums {
		numsMap[v] = append(numsMap[v], k)
	}
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			if nums[left] == nums[right] {
				return []int{numsMap[nums[left]][0], numsMap[nums[right]][1]}
			}
			return []int{numsMap[nums[left]][0], numsMap[nums[right]][0]}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{}
}
