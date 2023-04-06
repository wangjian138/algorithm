package main

func main() {

}

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}
