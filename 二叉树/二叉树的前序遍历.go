package main

func main() {

}

var res []int

func preorderTraversal(root *TreeNode) []int {
	preorder(root)
	return res
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	preorder(root.Left)
	preorder(root.Right)
}
