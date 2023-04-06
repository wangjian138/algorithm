package main

import "container/heap"

// https://leetcode.cn/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	// 小顶堆，堆顶是最小元素
	pq := priorityQueue{}
	for _, e := range nums {
		// 每个元素都要过一遍二叉堆
		pq.offer(e)
		// 堆中元素多于 k 个时，删除堆顶元素
		if pq.size() > k {
			pq.poll()
		}
	}
	// pq 中剩下的是 nums 中 k 个最大元素，
	// 堆顶是最小的那个，即第 k 个最大元素
	return pq.peek()
}

type priorityQueue []int

func (pq *priorityQueue) Len() int { return len(*pq) }

func (pq *priorityQueue) size() int { return len(*pq) }

func (pq *priorityQueue) Less(i, j int) bool { return (*pq)[i] < (*pq)[j] }

func (pq *priorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *priorityQueue) offer(e int) { heap.Push(pq, e) }

func (pq *priorityQueue) poll() int { return heap.Pop(pq).(int) }

func (pq *priorityQueue) peek() int { return (*pq)[0] }

// 详细解析参见：
// https://labuladong.github.io/article/?qno=215
