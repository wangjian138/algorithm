package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("d:%d\n", '1')
}

// https://leetcode.cn/problems/string-to-integer-atoi/?favorite=2ckc81c
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func myAtoi(str string) int {
	n := len(str)
	i := 0
	// 记录正负号
	sign := 1
	// 用 long 避免 int 溢出
	var res int64 = 0
	// 跳过前导空格
	for i < n && str[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}
	// 记录符号位
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}
	if i == n {
		return 0
	}
	// 统计数字位
	for i < n && '0' <= str[i] && str[i] <= '9' {
		res = res*10 + int64(str[i]-'0')
		if res > math.MaxInt32 {
			break
		}
		i++
	}
	// 如果溢出，强转成 int 就会和真实值不同
	if res > math.MaxInt32 {
		if sign == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	return int(res) * sign
}
