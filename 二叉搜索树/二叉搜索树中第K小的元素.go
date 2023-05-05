package main

func main() {

}

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// kthSmallest returns the kth smallest element in the BST rooted with root.
func kthSmallest(root *TreeNode, k int) int {
	// 记录结果
	var res int
	// 记录当前元素的排名
	var rank int

	// traverse recursively traverses the BST rooted with root in-order
	// and finds the rank-th smallest element in the BST.
	// It updates rank and res accordingly.
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		/* 中序遍历代码位置 */
		rank++
		if k == rank {
			// 找到第 k 小的元素
			res = root.Val
			return
		}
		/*****************/
		traverse(root.Right)
	}

	traverse(root)
	return res
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=230
