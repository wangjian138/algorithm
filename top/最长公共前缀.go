package main

func main() {

}

// https://leetcode.cn/problems/longest-common-prefix/?favorite=2ckc81c
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func longestCommonPrefix(strs []string) string {
	m := len(strs)
	// 以第一行的列数为基准
	n := len(strs[0])
	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			thisStr, prevStr := strs[row], strs[row-1]
			// 判断每个字符串的 col 索引是否都相同
			if col >= len(thisStr) || col >= len(prevStr) ||
				thisStr[col] != prevStr[col] {
				// 发现不匹配的字符，只有 strs[row][0..col-1] 是公共前缀
				return strs[row][:col]
			}
		}
	}
	return strs[0]
}
