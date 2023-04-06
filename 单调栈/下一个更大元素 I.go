package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(nums1, nums2)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	stack := make([]int, 0)

	nums2IncreaseMap := make(map[int]int, 0)
	for _, v := range nums2 {
		nums2IncreaseMap[v] = -1
	}

	for k, v := range nums2 {
		if k == 0 {
			stack = append(stack, 0)
			continue
		}
		for len(stack) > 0 && nums2[stack[len(stack)-1]] < v {
			top := stack[len(stack)-1]
			nums2IncreaseMap[nums2[top]] = v

			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, k)
	}
	for _, v := range nums1 {
		res = append(res, nums2IncreaseMap[v])
	}

	return res
}
