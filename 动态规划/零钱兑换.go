package main

import (
	"fmt"
	"math"
)

func main() {
	num := coinChange([]int{1, 2}, 3)
	fmt.Printf("coinChange num:%v\n", num)
}

// https://leetcode.cn/problems/coin-change/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = amount + 1
	}

	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange1(coins []int, amount int) int {
	coinsAll = make([]int, amount+1)

	for k := range coinsAll {
		coinsAll[k] = math.MinInt
	}

	return dp(coins, amount)
}

var coinsAll []int

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if coinsAll[amount] != math.MinInt {
		return coinsAll[amount]
	}
	res := math.MaxInt

	for _, coin := range coins {
		subRes := dp(coins, amount-coin)
		if subRes == -1 {
			continue
		}
		res = min(res, subRes+1)
	}
	if res == math.MaxInt {
		res = -1
	}

	coinsAll[amount] = res

	return res
}
