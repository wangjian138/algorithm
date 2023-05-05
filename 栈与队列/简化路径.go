package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/a//c/d"
	res := simplifyPath(path)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/simplify-path/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stk := make([]string, 0)
	// 借助栈计算最终的文件夹路径
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(stk) != 0 {
				stk = stk[:len(stk)-1]
			}
		} else {
			stk = append(stk, part)
		}
	}
	// 栈中存储的文件夹组成路径
	res := ""
	for i := len(stk) - 1; i >= 0; i-- {
		res = "/" + stk[i] + res
	}
	if res == "" {
		res = "/"
	}
	return res
}
