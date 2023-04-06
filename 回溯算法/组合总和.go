package main

func main() {

}

// https://leetcode.cn/problems/combination-sum/
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func combinationSum1(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	find(candidates, target, 0, []int{})
	return res
}

var res [][]int

func find(candidates []int, target, begin int, data []int) {
	if target == 0 {
		res = append(res, append([]int(nil), data...))
		return
	}

	for i := begin; i < len(candidates); i++ {
		v := candidates[i]
		if target-v < 0 {
			continue
		}
		data = append(data, v)
		find(candidates, target-v, i, data)
		data = data[0 : len(data)-1]
	}
}
