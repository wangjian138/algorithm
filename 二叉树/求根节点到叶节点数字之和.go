package main

func main() {

}

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/?envType=study-plan-v2&id=top-interview-150
func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

type pair struct {
	node *TreeNode
	num  int
}

func sumNumbers(root *TreeNode) (sum int) {
	if root == nil {
		return
	}
	queue := []pair{{root, root.Val}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		left, right, num := p.node.Left, p.node.Right, p.num
		if left == nil && right == nil {
			sum += num
		} else {
			if left != nil {
				queue = append(queue, pair{left, num*10 + left.Val})
			}
			if right != nil {
				queue = append(queue, pair{right, num*10 + right.Val})
			}
		}
	}
	return
}
