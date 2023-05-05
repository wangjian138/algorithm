package main

import "fmt"

func main() {
	fmt.Printf("num:%v\n", 3&2)
}

// https://leetcode.cn/problems/number-of-1-bits/?favorite=2ckc81c
func hammingWeight(n uint32) int {
	res := 0
	for n != 0 {
		n &= n - 1
		res++
	}
	return res
}
