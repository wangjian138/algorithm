package main

import (
	"fmt"
	"math"
)

func main() {
	profit := maxProfit([]int{1, 2, 3})

	fmt.Printf("prices: maxProfit:%v\n", profit)
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
func maxProfit(prices []int) int {
	dp := make(map[int][3][2]int, len(prices))

	for k, v := range prices {
		if k == 0 {
			dp[k] = [3][2]int{
				{0, math.MinInt},
				{0, -v},
				{0, -v},
			}
			continue
		}
		trade := dp[k]
		for i := 1; i < 3; i++ {
			trade[i] = [2]int{
				max(dp[k-1][i][0], dp[k-1][i][1]+v),
				max(dp[k-1][i][1], dp[k-1][i-1][0]-v),
			}
		}
		dp[k] = trade
	}
	tradeNum := 2
	if len(prices) <= 2 {
		tradeNum = 1
	}

	return dp[len(prices)-1][tradeNum][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
