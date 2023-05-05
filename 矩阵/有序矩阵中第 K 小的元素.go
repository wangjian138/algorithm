package main

import (
	"container/heap"
)

func main() {

}

// https://leetcode.cn/problems/kth-smallest-element-in-a-sorted-matrix/?show=1
func kthSmallest(matrix [][]int, k int) int {
	// 自定义一个最小堆类型
	pq := IntHeap{}
	// 初始化堆，把每一行的第一个元素装进去
	for i := 0; i < len(matrix); i++ {
		pq = append(pq, Item{value: matrix[i][0], row: i, col: 0})
	}
	heap.Init(&pq)

	var res int
	// 执行合并多个有序链表的逻辑，找到第 k 小的元素
	for k > 0 && pq.Len() > 0 {
		cur := heap.Pop(&pq).(Item)
		res = cur.value
		k--
		// 链表中的下一个节点加入堆
		row, col := cur.row, cur.col+1
		if col < len(matrix[row]) {
			heap.Push(&pq, Item{value: matrix[row][col], row: row, col: col})
		}
	}

	return res
}

// 定义一个 Item 类型，表示堆中的元素
type Item struct {
	value int // 当前元素的值
	row   int // 当前元素所在的行
	col   int // 当前元素所在的列
}

// 定义一个最小堆类型 IntHeap
// 实现 heap.Interface 接口的方法
type IntHeap []Item

func (t *IntHeap) Len() int {
	return len(*t)
}

func (t *IntHeap) Less(i, j int) bool {
	return (*t)[i].value < (*t)[j].value
}

func (t *IntHeap) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *IntHeap) Push(x interface{}) {
	*t = append(*t, x.(Item))
}

func (t *IntHeap) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}
