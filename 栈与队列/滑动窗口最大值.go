package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 5, 1, 4}
	//nums := []int{1}
	k := 3
	arr := maxSlidingWindow(nums, k)
	fmt.Printf("arr:%v\n", arr)
}

type MyQueue struct {
	nums []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		nums: make([]int, 0),
	}
}

func (q *MyQueue) Front() int {
	return q.nums[0]
}

func (q *MyQueue) Back() int {
	return q.nums[len(q.nums)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.nums) == 0
}

func (q *MyQueue) Push(val int) {
	for !q.Empty() && val > q.Back() {
		q.nums = q.nums[:len(q.nums)-1]
	}

	q.nums = append(q.nums, val)
}

func (q *MyQueue) Pop(val int) {
	if !q.Empty() && val == q.Front() {
		q.nums = q.nums[1:]
	}
}

// https://leetcode.cn/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	res := make([]int, 0)

	// 加入前面k个元素
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	numsLen := len(nums)

	// 添加前面k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < numsLen; i++ {
		// 删除原来第i-k个元素
		queue.Pop(nums[i-k])

		// 添加当前的元素
		queue.Push(nums[i])

		// 获取最大值
		res = append(res, queue.Front())
	}

	return res
}
