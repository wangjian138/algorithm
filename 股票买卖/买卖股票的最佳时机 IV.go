package main

import (
	"fmt"
	"math"
)

func main() {
	profit := maxProfit(1, []int{3, 4})
	fmt.Printf("profit:%v\n", profit)
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
func maxProfit(k int, prices []int) int {
	dp := make([][][2]int, len(prices))
	dpk := make([][2]int, k+1)

	for i, v := range prices {
		if i == 0 {
			for j := 0; j <= k; j++ {
				if j == 0 {
					dp[i] = append(dp[i], [2]int{0, math.MinInt})
					continue
				}
				dp[i] = append(dp[i], [2]int{0, -v})
			}
			continue
		}
		for j := 1; j <= k; j++ {
			dpk[j] = [2]int{
				max(dp[i-1][j][0], dp[i-1][j][1]+v),
				max(dp[i-1][j][1], dp[i-1][j-1][0]-v),
			}
		}

		dp[i] = dpk
	}
	fmt.Printf("dp:%v\n", dp)
	return dp[len(prices)-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
