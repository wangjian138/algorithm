package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/h-index/
func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}
	return
}
