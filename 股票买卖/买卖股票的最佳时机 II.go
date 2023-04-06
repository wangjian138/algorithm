package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	dp := make(map[int][2]int, 0)

	pricesLen := len(prices)

	for k, v := range prices {
		if k == 0 {
			dp[k] = [2]int{
				0, -v,
			}
			continue
		}
		dp[k] = [2]int{
			max(dp[k-1][0], dp[k-1][1]+v),
			max(dp[k-1][1], dp[k-1][0]-v),
		}
	}

	return dp[pricesLen-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
