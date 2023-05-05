package main

import "fmt"

func main() {
	str := "1a"
	num := lengthOfLastWord(str)
	fmt.Printf("num:%v\n", num)
}

// https://leetcode.cn/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return sLen
	}

	index := sLen - 1
	for index >= 0 && s[index] == ' ' {
		index--
	}
	s = s[:index+1]

	num := 0
	index = len(s) - 1
	for index >= 0 && s[index] != ' ' {
		num++
		index--
	}
	return num
}
