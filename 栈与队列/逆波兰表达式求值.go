package main

import (
	"fmt"
	"strconv"
)

func main() {
	token := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	res := evalRPN(token)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, b := range tokens {
		num, err := strconv.Atoi(b)
		if err == nil {
			stack = append(stack, num)
			continue
		}
		num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[0 : len(stack)-2]
		if b == "+" {
			stack = append(stack, num1+num2)
		} else if b == "-" {
			stack = append(stack, num1-num2)
		} else if b == "*" {
			stack = append(stack, num1*num2)
		} else if b == "/" {
			stack = append(stack, num1/num2)
		}
	}
	return stack[len(stack)-1]
}
