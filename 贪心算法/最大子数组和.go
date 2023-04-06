package main

func main() {

}

// https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	dp := make([]int, numsLen)

	dp[0] = nums[0]
	var res int
	res = dp[0]
	for i := 1; i < numsLen; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
