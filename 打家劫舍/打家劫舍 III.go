package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := getRob(root)
	return max(res[0], res[1])
}

func getRob(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}
	left := getRob(root.Left)
	right := getRob(root.Right)
	robRoot := root.Val + left[0] + right[0]
	noRobRoot := left[1] + right[1]

	return [2]int{noRobRoot, robRoot}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
