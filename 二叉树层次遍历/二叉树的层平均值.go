package main

func main() {

}

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/
func averageOfLevels(root *TreeNode) []float64 {
	var levelData []data
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].count++
			levelData[level].val += node.Val
		} else {
			levelData = append(levelData, data{
				count: 1,
				val:   node.Val,
			})
		}
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)

	res := make([]float64, len(levelData))

	for i, v := range levelData {
		res[i] = float64(v.val) / float64(v.count)
	}
	return res
}

type data struct {
	val, count int
}
