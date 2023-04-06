package main

import "sort"

func main() {

}

// https://leetcode.cn/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	// 从小到大
	child := 0
	for sIdx := 0; child < len(g) && sIdx < len(s); sIdx++ {
		if s[sIdx] >= g[child] { //如果饼干的大小大于或等于孩子的为空则给与，否则不给予，继续寻找选一个饼干是否符合
			child++
		}
	}

	return child
}
