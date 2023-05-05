package main

import "math"

func main() {

}

// https://leetcode.cn/problems/minimum-path-sum/?envType=study-plan-v2
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 构造备忘录，初始值全部设为 -1
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		if i < 0 || j < 0 {
			return int(math.MaxInt64)
		}
		// 避免重复计算
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		// 将计算结果记入备忘录
		memo[i][j] = min(
			dp(i-1, j),
			dp(i, j-1),
		) + grid[i][j]

		return memo[i][j]
	}
	return dp(m-1, n-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=64
