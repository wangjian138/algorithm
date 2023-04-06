package main

func main() {

}

// https://leetcode.cn/problems/subsets/
func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	all(nums, 0, []int{})
	return res
}

var res [][]int

func all(nums []int, begin int, data []int) {
	res = append(res, append([]int(nil), data...))
	for i := begin; i < len(nums); i++ {
		data = append(data, nums[i])
		all(nums, i+1, data)
		data = data[0 : len(data)-1]
	}
}
