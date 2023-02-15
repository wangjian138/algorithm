package main

import "fmt"

func main() {
	t := &TreeNode{
		Val:   1,
		Right: nil,
		Left: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(maxDepth(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	traverse(root)
	return res
}

var res, depth int

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		if res < depth {
			res = depth
		}
	}
	traverse(root.Left)
	traverse(root.Right)
	depth--
}

func maxDepth1(root *TreeNode) int {

	var left int
	left = 0

	maxNum := getreeLenMax(root, left)

	return maxNum
}

func getreeLenMax(node *TreeNode, num int) (maxNum int) {
	if node == nil {
		maxNum = num
		return
	}

	if node.Left == nil && node.Right == nil {
		num++
		maxNum = num
		return
	}

	if node.Left != nil && node.Right != nil {
		num++
		maxNumLeft := getreeLenMax(node.Left, num)
		maxNumRight := getreeLenMax(node.Right, num)
		if maxNumLeft > maxNumRight {
			maxNum = maxNumLeft
		} else {
			maxNum = maxNumRight
		}
		return
	} else if node.Left != nil {
		num++
		maxNum = getreeLenMax(node.Left, num)
	} else if node.Right != nil {
		num++
		maxNum = getreeLenMax(node.Right, num)
	}

	return
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
