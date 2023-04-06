package main

import (
	"fmt"
	"math"
)

func main() {
	x := -1123
	res := reverse(x)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/reverse-integer/
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
		//fmt.Printf("digit:%v x:%v rev:%v\n", digit, x, rev)
	}
	return
}
