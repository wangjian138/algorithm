package main

import "fmt"

func main() {

	testNum := 2
	res1 := isHappy(testNum)
	fmt.Printf("res:%v\n", res1)
}

// https://leetcode.cn/problems/happy-number/
func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	fmt.Printf("fast:%v\n", fast)
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
