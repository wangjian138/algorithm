package main

func main() {

}

// https://leetcode.cn/problems/word-break/?favorite=2ckc81c
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。
func wordBreak(s string, wordDict []string) bool {
	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}
	return dp(s, 0, wordDict, memo)
}

// dp：返回 s[i..] 是否能够被 wordDict 拼出
func dp(s string, i int, wordDict []string, memo []int) bool {
	// base case: 整个 s 都被拼出来了
	if i == len(s) {
		return true
	}
	// 防止冗余计算
	if memo[i] != -1 {
		return memo[i] == 1
	}
	// 遍历所有单词，尝试匹配 s[i..] 的前缀
	for _, word := range wordDict {
		len1 := len(word)
		if i+len1 > len(s) {
			continue
		}
		subStr := s[i : i+len1]
		if subStr != word {
			continue
		}
		// s[i..] 的前缀被匹配，去尝试匹配 s[i+len..]
		if dp(s, i+len1, wordDict, memo) {
			// s[i..] 可以被拼出，将结果记入备忘录
			memo[i] = 1
			return true
		}
	}
	// s[i..] 不能被拼出，结果记入备忘录
	memo[i] = 0
	return false
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=139
