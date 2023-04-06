package main

import "fmt"

func main() {
	robNum := []int{9, 2, 7, 10}
	robRes := rob(robNum)
	fmt.Println("res:", robRes)
}

// https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
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
