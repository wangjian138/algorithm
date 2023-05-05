package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	res3 = make([][]int, 0)
	sort.Ints(nums)
	getAll(nums, []int{}, map[int]bool{})
	return res3
}

var res3 [][]int

func getAll(nums []int, data []int, filter map[int]bool) {
	if len(data) == len(nums) {
		res3 = append(res3, append([]int(nil), data...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if filter[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !filter[i-1] {
			continue
		}
		filter[i] = true
		data = append(data, nums[i])
		getAll(nums, data, filter)
		data = data[0 : len(data)-1]
		filter[i] = false
	}
}
