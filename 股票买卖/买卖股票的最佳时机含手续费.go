package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
func maxProfit(prices []int, fee int) int {
	dp := make(map[int][2]int, len(prices))

	dp[-1] = [2]int{
		0,
		-1 << 20,
	}
	for k, v := range prices {
		dp[k] = [2]int{
			max(dp[k-1][0], dp[k-1][1]+v-fee),
			max(dp[k-1][1], dp[k-1][0]-v),
		}
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
