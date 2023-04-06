package main

func main() {

}

// https://leetcode.cn/problems/combinations/
func combine(n int, k int) [][]int {
	res = make([][]int, 0)
	backtrack(n, k, 1, []int{})
	return res
}

var res [][]int

func backtrack(n int, k, begin int, data []int) {
	if len(data) == k {
		res = append(res, append([]int(nil), data...))
		return
	}

	for i := begin; i <= n; i++ {
		data = append(data, i)
		backtrack(n, k, i+1, data)
		data = data[0 : len(data)-1]
	}
}
