package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	fmt.Printf("nums:%v\n", res)
}

// https://leetcode.cn/problems/longest-increasing-subsequence/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

func lengthOfLIS(nums []int) int {
	// dp[i] 表示以 nums[i] 这个数结尾的最长递增子序列的长度
	dp := make([]int, len(nums))
	// base case：dp 数组全都初始化为 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := 0
	for i := 0; i < len(dp); i++ {
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

// 详细解析参见：
// https://labuladong.github.io/article/?qno=300
