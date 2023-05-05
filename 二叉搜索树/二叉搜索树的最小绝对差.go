package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码还未经过力扣测试，仅供参考，如有疑惑，可以参照我写的 java 代码对比查看。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	prev := (*TreeNode)(nil)
	res := math.MaxInt32

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		// 中序遍历位置
		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}

		prev = root
		traverse(root.Right)
	}

	traverse(root)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
