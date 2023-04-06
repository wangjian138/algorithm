package main

func main() {

}

// https://leetcode.cn/problems/longest-valid-parentheses/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func longestValidParentheses(s string) int {
	stk := make([]int, 0)
	// dp[i] 的定义：记录以 s[i-1] 结尾的最长合法括号子串长度
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 遇到左括号，记录索引
			stk = append(stk, i)
			// 左括号不可能是合法括号子串的结尾
			dp[i+1] = 0
		} else {
			// 遇到右括号
			if len(stk) != 0 {
				// 配对的左括号对应索引
				leftIndex := stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				// 以这个右括号结尾的最长子串长度
				len := 1 + i - leftIndex + dp[leftIndex]
				dp[i+1] = len
			} else {
				// 没有配对的左括号
				dp[i+1] = 0
			}
		}
	}
	// 计算最长子串的长度
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

func longestValidParentheses1(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}
