package main

import "fmt"

func main() {
	res := jump1([]int{2, 3, 1, 1, 4})
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	var res int
	numLen := len(nums)
	if numLen == 0 {
		return 0
	}

	var stack []int
	stack = append(stack, 0)

	var maxNum int
	var maxIndex int
	for len(stack) != 0 {
		res++
		firstNum := nums[stack[0]]
		firstIndex := stack[0]
		maxNum = 0
		if firstIndex >= numLen-1 {
			return res
		}

		for i := 1; i <= firstNum; i++ {
			nextIndex := firstIndex + i
			currentNum := nums[nextIndex] + nextIndex
			if currentNum > maxNum {
				maxNum = currentNum
				maxIndex = nextIndex
			}
		}
		if nums[maxIndex] == 0 {
			return 0
		}
		stack = []int{maxIndex}
	}

	return maxIndex
}

func jump1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	cur, next := 0, 0
	step := 0
	for i := 0; i < n-1; i++ {
		next = max(nums[i]+i, next)
		if i == cur {
			cur = next
			step++
		}
	}
	return step
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
