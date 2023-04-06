package main

import "fmt"

func main() {
	res := climbStairs(3)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	var dp1, dp2 int
	dp1 = 1
	dp2 = 2
	for i := 2; i < n; i++ {
		temp := dp2
		dp2 = dp1 + dp2 + 1
		dp1 = temp
	}
	return dp2
}
