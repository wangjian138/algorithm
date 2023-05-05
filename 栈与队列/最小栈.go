package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(-3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	res := stack.GetMin()
	fmt.Printf("res:%v\n", res)

	stack.Pop()

	res = stack.GetMin()
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/min-stack/?envType=study-plan-v2&id=top-interview-150
type MinStack struct {
	// 记录栈中的所有元素
	stk []int
	// 阶段性记录栈中的最小元素
	minStk []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	// 维护 minStk 栈顶为全栈最小元素
	if len(this.minStk) == 0 || val <= this.minStk[len(this.minStk)-1] {
		// 新插入的这个元素就是全栈最小的
		this.minStk = append(this.minStk, val)
	}
}

func (this *MinStack) Pop() {
	// 注意 Go 语言的语言特性，比较 int 相等直接用 ==
	if this.stk[len(this.stk)-1] == this.minStk[len(this.minStk)-1] {
		// 弹出的元素是全栈最小的
		this.minStk = this.minStk[:len(this.minStk)-1]
	}
	this.stk = this.stk[:len(this.stk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	// minStk 栈顶为全栈最小元素
	return this.minStk[len(this.minStk)-1]
}
