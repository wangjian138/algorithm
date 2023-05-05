package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	intervalsLen := len(intervals)
	if intervalsLen <= 1 {
		return intervals
	}

	sort.Sort(twoSortInts(intervals))

	left := intervals[0][0]
	right := intervals[0][1]

	var intervalsNew [][]int

	intervalsNew = append(intervalsNew, []int{left, right})

	for i := 1; i < intervalsLen; i++ {
		if intervals[i][0] > right {
			left = intervals[i][0]
			right = intervals[i][1]
			intervalsNew = append(intervalsNew, []int{left, right})
			continue
		}

		if intervals[i][0] >= left && intervals[i][1] >= right {
			right = intervals[i][1]
			intervalsNew[len(intervalsNew)-1][1] = right
			continue
		}
	}

	return intervalsNew
}

type twoSortInts [][]int

func (t twoSortInts) Len() int {
	return len(t)
}

func (t twoSortInts) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t twoSortInts) Less(i, j int) bool {
	if t[i][0] < t[j][0] {
		return true
	}

	if t[i][0] == t[j][0] && t[i][1] >= t[j][1] {
		return true
	}

	return false
}
