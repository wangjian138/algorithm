package main

func main() {

}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
	dp := make([][2]int, 0)

	for k, v := range prices {
		if k == 0 {
			dp[0] = [2]int{
				0, -v,
			}
			continue
		}
		if k == 1 {
			dp[1] = [2]int{
				max(dp[0][0], dp[0][1]+v),
				max(dp[0][1], dp[0][0]-v),
			}
			continue
		}

		dp[k][0] = max(dp[k-1][0], dp[k-1][1]+v)
		dp[k][1] = max(dp[k-1][1], dp[k-2][0]-v)
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
