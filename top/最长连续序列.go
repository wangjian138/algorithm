package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 0, 1}
	res := longestConsecutive(nums)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/longest-consecutive-sequence/?favorite=2ckc81c
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func longestConsecutive(nums []int) int {
	// 转化成哈希集合，方便快速查找是否存在某个元素
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	res := 0

	for num := range set {
		if set[num-1] {
			// num 不是连续子序列的第一个，跳过
			continue
		}
		// num 是连续子序列的第一个，开始向上计算连续子序列的长度
		curNum := num
		curLen := 1

		for set[curNum+1] {
			curNum += 1
			curLen += 1
		}
		// 更新最长连续序列的长度
		res = max(res, curLen)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
