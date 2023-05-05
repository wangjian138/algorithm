package main

func main() {

}

// https://leetcode.cn/problems/same-tree/?envType=study-plan-v2&id=top-interview-150
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := isSameTree(p.Left, q.Left)
	if !left {
		return false
	}
	right := isSameTree(p.Right, q.Right)
	if !right {
		return false
	}
	return true
}
