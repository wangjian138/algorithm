package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 5},
		{6, 8},
	}
	newInterval := []int{0, 9}
	res := insert(intervals, newInterval)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/insert-interval/?envType=study-plan-v2&id=top-interview-150
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	intervals = append(intervals, newInterval)
	intervalsLen := len(intervals)

	sort.Sort(twoSlice(intervals))

	first := intervals[0][0]
	second := intervals[0][1]

	res := make([][]int, 0)
	res = append(res, []int{first, second})

	for i := 1; i < intervalsLen; i++ {
		if res[len(res)-1][1] == intervals[i][0] {
			res[len(res)-1][1] = intervals[i][1]
			continue
		}

		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
			continue
		}

		if res[len(res)-1][1] < intervals[i][1] {
			res[len(res)-1][1] = intervals[i][1]
		}
	}

	return res
}

type twoSlice [][]int

func (t twoSlice) Len() int {
	return len(t)
}

func (t twoSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t twoSlice) Less(i, j int) bool {
	if t[i][0] <= t[j][0] {
		return true
	}

	return false
}
