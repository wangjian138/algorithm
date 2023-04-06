package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	res2 = make([][]int, 0)
	sort.Ints(nums)
	all2(nums, 0, []int{})
	return res2
}

var res2 [][]int

func all2(nums []int, begin int, data []int) {
	res2 = append(res2, append([]int(nil), data...))

	for i := begin; i < len(nums); i++ {
		if i > begin && nums[i] == nums[i-1] {
			continue
		}
		data = append(data, nums[i])
		all2(nums, i+1, data)
		data = data[0 : len(data)-1]
	}
}
