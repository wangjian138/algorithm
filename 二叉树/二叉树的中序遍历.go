package main

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	traversal(root)
	return res
}

var res []int

func traversal(root *TreeNode) {
	if root == nil {
		return
	}

	traversal(root.Left)
	res = append(res, root.Val)
	traversal(root.Right)
}
