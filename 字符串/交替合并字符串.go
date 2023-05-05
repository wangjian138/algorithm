package main

import "fmt"

func main() {
	s1 := "abc1"
	s2 := "pqr2345"
	str := mergeAlternately(s1, s2)
	fmt.Printf("str:%v\n", str)
}

// https://leetcode.cn/problems/merge-strings-alternately/?envType=study-plan-v2&id=leetcode-75
func mergeAlternately(word1 string, word2 string) string {
	w1Len := len(word1)
	w2len := len(word2)

	newStr := make([]byte, 0)
	index := 0
	for index < w1Len && index < w2len {
		newStr = append(newStr, word1[index], word2[index])
		index++
	}
	if w1Len > index {
		newStr = append(newStr, word1[index:]...)
	} else {
		newStr = append(newStr, word2[index:]...)

	}
	return string(newStr)
}
