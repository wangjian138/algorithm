package main

func main() {

}

// https://leetcode.cn/problems/number-of-enclaves/
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m-1, j)
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += 1
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=1020
