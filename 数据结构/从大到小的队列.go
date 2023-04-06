package main

func main() {

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
