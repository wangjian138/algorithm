package main

import "fmt"

func main() {
	s := "aba"
	fmt.Printf("encode:%v\n", encode(s))
}

// https://leetcode.cn/problems/group-anagrams/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func groupAnagrams(strs []string) [][]string {
	// 编码到分组的映射
	codeToGroup := make(map[string][]string)
	for _, s := range strs {
		// 对字符串进行编码
		code := encode(s)
		// 把编码相同的字符串放在一起
		codeToGroup[code] = append(codeToGroup[code], s)
	}

	// 获取结果
	res := make([][]string, 0, len(codeToGroup))
	for _, group := range codeToGroup {
		res = append(res, group)
	}

	return res
}

// 利用每个字符的出现次数进行编码
func encode(s string) string {
	count := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		delta := s[i] - 'a'
		count[delta]++
	}
	return string(count)
}
