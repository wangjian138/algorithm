package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%b\n", -1)
	prices := []int{1, 2, 3, 4, 2, 1}
	num := maxProfit(prices)
	fmt.Printf("num:%v\n", num)
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	dp := make(map[int][2]int, len(prices))

	dp[-1] = [2]int{0, math.MinInt}

	for k := range prices {
		dp[k] = [2]int{getMax(dp[k-1][0], dp[k-1][1]+prices[k]), getMax(dp[k-1][1], -prices[k])}
	}

	return dp[len(prices)-1][0]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfitCompress(prices []int) int {
	n := len(prices)
	// base case: dp[-1][0] = 0, dp[-1][1] = -infinity
	dp_i_0, dp_i_1 := 0, -2147483648
	for i := 0; i < n; i++ {
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		// dp[i][1] = max(dp[i-1][1], -prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
