package main

import "fmt"

func main() {
	node := &TreeNode{Left: &TreeNode{Right: &TreeNode{Left: &TreeNode{}}}}

	depth := minDepth(node)
	fmt.Printf("depth:%v\n", depth)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	var treeNodeList []*TreeNode
	treeNodeList = append(treeNodeList, root)

	for len(treeNodeList) > 0 {
		treeNodeListLen := len(treeNodeList)
		for i := 0; i < treeNodeListLen; i++ {
			node := treeNodeList[0]
			treeNodeList = treeNodeList[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				treeNodeList = append(treeNodeList, node.Left)
			}

			if node.Right != nil {
				treeNodeList = append(treeNodeList, node.Right)
			}
		}
		depth++
	}

	return depth
}
