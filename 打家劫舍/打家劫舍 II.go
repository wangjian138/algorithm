package main

import (
	"fmt"
)

func main() {
	robNum := []int{1, 2, 1}
	robRes := rob(robNum)
	fmt.Println("robRes:", robRes)
}

// https://leetcode.cn/problems/house-robber-ii/
func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	} else if numsLen == 1 {
		return nums[0]
	} else if numsLen == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob1(nums[0:numsLen-1]), rob1(nums[1:]))
}

func rob1(nums []int) int {
	dp := make([][2]int, len(nums))

	for k, v := range nums {
		if k == 0 {
			dp[k] = [2]int{
				0,
				v,
			}
			continue
		}
		if k == 1 {
			dp[k] = [2]int{dp[k-1][1], v}
			continue
		}
		dp[k] = [2]int{max(dp[k-1][1], dp[k-1][0]), max(dp[k-2][1]+v, dp[k-2][0]+v)}
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
