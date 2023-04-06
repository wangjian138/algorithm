package main

import (
	"fmt"
	"sort"
)

func main() {
	heap1 := &IHeap{2, 4, 1, 6}
	//heap1 := &IHeap{}
	Init(heap1)
	//Push(heap1, 2)
	//Push(heap1, 7)
	//Push(heap1, 5)
	//Push(heap1, 5)
	//Push(heap1, 4)
	//Push(heap1, 6)
	fmt.Printf("heap1:%v\n", heap1)
	for len(*heap1) != 0 {
		fmt.Println(Pop(heap1))
	}
}

// 构建小顶堆 这个是需要实现的interface
type IHeap []int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x any) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		println("step1 j1:", j1)
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		println("step2 j:", j)
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		println("step3 j:", j)
		if !h.Less(j, i) {
			println("break")
			break
		}
		println("swap i:", i, " j:", j)
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
