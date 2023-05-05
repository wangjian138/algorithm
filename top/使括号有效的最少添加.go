package main

func main() {

}

// https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
func minAddToMakeValid(s string) int {
	if len(s) == 0 {
		return 0
	}

	stack := make([]byte, 0)

	for _, b := range s {
		if len(stack) == 0 || b == '(' {
			stack = append(stack, byte(b))
			continue
		}
		// 这个时候就是 )
		lastStr := stack[len(stack)-1]
		if lastStr == ')' {
			stack = append(stack, byte(b))
			continue
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
