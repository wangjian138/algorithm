package main

import "fmt"

func main() {
	res := generate(3)
	fmt.Printf("res:%v\n", res)
}

// https://leetcode.cn/problems/pascals-triangle/
func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
				continue
			}
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}

	return res
}
