package main

func main() {

}

// https://leetcode.cn/problems/fibonacci-number/
func fib(n int) int {
	if n < 2 {
		return n
	}

	i_0 := 0
	i_1 := 1
	for i := 2; i <= n; i++ {
		temp := i_1
		i_1 = i_1 + i_0
		i_0 = temp
	}

	return i_1
}
