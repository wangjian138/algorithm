package main

func main() {

}

// https://leetcode.cn/problems/unique-paths-ii/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

// 第一步：自顶向下带备忘录的递归
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	return dp(obstacleGrid, m-1, n-1, memo)
}

// 定义：从 grid[0][0] 出发到达 grid[i][j] 的路径条数为 dp(grid, i, j)
func dp(grid [][]int, i int, j int, memo [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// base case
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 1 {
		// 数组越界或者遇到阻碍
		return 0
	}
	if i == 0 && j == 0 {
		// 起点到起点的路径条数就是 1
		return 1
	}
	if memo[i][j] > 0 {
		// 避免冗余计算
		return memo[i][j]
	}
	// 状态转移方程：
	// 到达 grid[i][j] 的路径条数等于
	// 到达 grid[i-1][j] 的路径条数加上到达 grid[i][j-1] 的路径条数
	left := dp(grid, i, j-1, memo)
	upper := dp(grid, i-1, j, memo)
	res := left + upper
	// 存储备忘录
	memo[i][j] = res
	return res
}

// 第二步：自底向上迭代的动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	// 数组索引偏移一位，dp[0][..] dp[..][0] 代表 obstacleGrid 之外
	// 定义：到达 obstacleGrid[i][j] 的路径条数为 dp[i-1][j-1]
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// base case：如果没有障碍物，起点到起点的路径条数就是 1
	dp[1][1] = 0
	if obstacleGrid[0][0] == 0 {
		dp[1][1] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				// 跳过 base case
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				// 跳过障碍物的格子
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	// 返回到达 obstacleGrid[m-1][n-1] 的路径数量
	return dp[m][n]
}

// 第三步：优化二维 dp 数组为一维 dp 数组
func uniquePathsWithObstacles3(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	// 根据二维 dp 数组优化为一维 dp 数组
	dp := make([]int, n+1)
	dp[1] = 0
	if obstacleGrid[0][0] == 0 {
		dp[1] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				// 跳过 base case
				continue
			}
			if obstacleGrid[i-1][j-1] == 1 {
				// 跳过障碍物的格子
				dp[j] = 0
				continue
			}
			dp[j] = dp[j] + dp[j-1]
		}
	}
	// 返回到达 obstacleGrid[m-1][n-1] 的路径数量
	return dp[n]
}
