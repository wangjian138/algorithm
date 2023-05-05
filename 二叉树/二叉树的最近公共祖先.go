package main

func main() {

}

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/?envType=study-plan-v2&id=top-interview-150
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 情况 1
	if left != nil && right != nil {
		return root
	}
	// 情况 2
	if left == nil && right == nil {
		return nil
	}
	// 情况 3
	if left == nil {
		return right
	}
	return left
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=236
