package main

import "fmt"

func main() {
	singleStr := removeDuplicates("abccd")
	fmt.Printf("singleStr:%v\n", singleStr)
}

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		// 栈不空 且 与栈顶元素不等
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// 弹出栈顶元素 并 忽略当前元素(s[i])
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
