package main

import "fmt"

func main() {
	testNums := []int{0, 0, 0, 1, 2, 3}
	leftNum := findLeft(testNums, 0)
	fmt.Printf("leftNum:%v\n", leftNum)

	rightNum := findRight(testNums, 0)
	fmt.Printf("rightNum:%v\n", rightNum)
}

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{findLeft(nums, target), findRight(nums, target)}
}

func findLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left < 0 || left >= len(nums) {
		return -1
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func findRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	if right < 0 {
		return -1
	}

	if nums[right] == target {
		return right
	}

	return -1
}
