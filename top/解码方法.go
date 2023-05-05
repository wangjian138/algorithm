package main

import "fmt"

func main() {
	s := numDecodings("111")
	fmt.Printf("s:%v\n", s)
}

// https://leetcode.cn/problems/decode-ways/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func numDecodings(s string) int {
	n := len(s)
	if n < 1 {
		return 0
	}
	// 定义：dp[i] 表示 s[0..i-1] 的解码方式数量
	dp := make([]int, n+1)
	// base case: s 为空或者 s 只有一个字符的情况
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	// 注意 dp 数组和 s 之间的索引偏移一位
	for i := 2; i <= n; i++ {
		c := s[i-1]
		d := s[i-2]
		if '1' <= c && c <= '9' {
			// 1. s[i] 本身可以作为一个字母
			dp[i] += dp[i-1]
		}
		fmt.Printf("i:%v dp:%v\n", i, dp)

		if d == '1' || d == '2' && c <= '6' {
			// 2. s[i] 和 s[i - 1] 结合起来表示一个字母
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
