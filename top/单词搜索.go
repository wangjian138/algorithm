package main

func main() {

}

// https://leetcode.cn/problems/word-search/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	found := false
	var dfs func(i, j, p int)
	dfs = func(i, j, p int) {
		if p == len(word) {
			// 整个 word 已经被匹配完，找到了一个答案
			found = true
			return
		}
		if found {
			// 已经找到了一个答案，不用再搜索了
			return
		}
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if board[i][j] != word[p] {
			return
		}

		// 已经匹配过的字符，我们给它添一个负号作为标记，避免走回头路
		board[i][j] ^= 255
		// word[p] 被 board[i][j] 匹配，开始向四周搜索 word[p+1..]
		dfs(i+1, j, p+1)
		dfs(i, j+1, p+1)
		dfs(i-1, j, p+1)
		dfs(i, j-1, p+1)
		board[i][j] ^= 255
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 0)
			if found {
				return true
			}
		}
	}
	return false
}
