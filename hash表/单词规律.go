package main

import "strings"

func main() {

}

// https://leetcode.cn/problems/word-pattern/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	// 记录 pattern 字符到单词的映射
	patternToWord := make(map[byte]string)
	// 记录那些已经有 pattern 对应的单词
	wordSet := make(map[string]bool)

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		word := words[i]
		if _, ok := patternToWord[c]; !ok {
			if _, ok := wordSet[word]; ok {
				// 这个单词以前已经有其他模式字符对应了
				return false
			}
			// 添加 c -> word 的映射
			patternToWord[c] = word
		} else {
			// 这个 pattern 字符已经出现过，确保和之前对应的单词相同
			if patternToWord[c] != word {
				return false
			}
		}
		// 这个单词已经有模式字符对应
		wordSet[word] = true
	}
	return true
}

func wordPattern2(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	// 记录 word 到 pattern 字符的映射
	wordToPattern := make(map[string]byte)
	// 记录那些已经有单词对应的 pattern 字符
	patternCharSet := make(map[byte]bool)

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		word := words[i]
		if _, ok := wordToPattern[word]; !ok {
			// 当前这个单词还没有对应的模式字符
			if _, ok := patternCharSet[c]; ok {
				// 对应的模式字符之前已经对应了其他单词
				return false
			}
			// 添加 word -> c 的映射
			wordToPattern[word] = c
		} else {
			// 这个单词之前已经出现过，确保当前单词和之前对应的模式字符相同
			if wordToPattern[word] != c {
				return false
			}
		}
		patternCharSet[c] = true
	}
	return true
}
