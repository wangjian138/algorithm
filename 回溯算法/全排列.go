package main

import "fmt"

func main() {

	res := permute([]int{5, 4, 6, 2})
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	tree = make([][]int, 0)
	find := make(map[int]bool)

	findRes(nums, []int{}, find)
	return tree
}
func findRes(nums []int, res []int, used map[int]bool) {
	if len(res) == len(nums) {
		newRes = make([]int, len(res))
		copy(newRes, res)
		tree = append(tree, newRes)
		return
	}

	for k, v := range nums {
		if used[k] {
			continue
		}
		fmt.Printf("findRes2 res:%v v:%v\n", res, v)
		used[k] = true
		res = append(res, v)
		findRes(nums, res, used)
		used[k] = false
		res = res[0 : len(res)-1]
		fmt.Printf("findRes3 res:%v v:%v\n", res, v)
	}
}

var tree [][]int
var newRes []int
