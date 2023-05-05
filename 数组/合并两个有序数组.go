package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/merge-sorted-array/?favorite=2ckc81c
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
