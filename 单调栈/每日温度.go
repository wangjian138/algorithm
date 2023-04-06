package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures1(temperatures)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	t := temperatures
	var res []int
	for i := 0; i < len(t)-1; i++ {
		j := i + 1
		for ; j < len(t); j++ {
			// 如果之后出现更高，说明找到了
			if t[j] > t[i] {
				res = append(res, j-i)
				break
			}
		}
		// 找完了都没找到
		if j == len(t) {
			res = append(res, 0)
		}
	}
	// 最后一个肯定是 0
	return append(res, 0)
}

func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, 0)

	stack := make([]int, 0)

	for k, v := range temperatures {
		stackLen := len(stack)
		if k == 0 || stackLen == 0 {
			stack = append(stack, v)
			continue
		}
		if v == stack[stackLen-1] || v < stack[stackLen-1] {
			stack = append(stack, v)
			continue
		}

		fmt.Printf("start v:%v stack:%v res:%v\n", v, stack, res)
		var j int
		for j = stackLen - 1; j >= 0; j-- {
			fmt.Printf("v:%v stack[j]:%v k:%v j:%v\n", v, stack[j], k, j)
			if stack[j] >= v {
				break
			}
			res = append(res, k-j)
		}
		stack = stack[0 : j+1]
		stack = append(stack, v)
	}
	return res
}

func dailyTemperatures2(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 初始化栈顶元素为第一个下标索引0
	stack := []int{0}

	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}
	return res
}
