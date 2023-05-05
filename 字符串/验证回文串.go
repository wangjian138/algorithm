package main

import (
	"fmt"
)

func main() {
	str := "race a car"
	res := isPalindrome(str)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/valid-palindrome/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。
func isPalindrome(s string) bool {
	// 先把所有字符转化成小写，并过滤掉空格和标点这类字符
	newStr := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			newStr = append(newStr, c)
			continue
		}
		if c >= 'A' && c <= 'Z' {
			newStr = append(newStr, c+'a'-'A')
			continue
		}
	}

	// 然后对剩下的这些目标字符执行双指针算法，判断回文串
	// 一左一右两个指针相向而行
	left, right := 0, len(newStr)-1
	for left < right {
		if newStr[left] != newStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
