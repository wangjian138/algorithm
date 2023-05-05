package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		temp := make([]int, 0)
		listTemp := make([]*TreeNode, 0)
		for i := 0; i < len(list); i++ {
			if list[i].Right != nil {
				listTemp = append(listTemp, list[i].Right)
				if len(temp) == 0 {
					temp = append(temp, list[i].Right.Val)
				}
			}
			if list[i].Left != nil {
				listTemp = append(listTemp, list[i].Left)
				if len(temp) == 0 {
					temp = append(temp, list[i].Left.Val)
				}
			}
		}

		res = append(res, temp...)
		list = listTemp
	}

	return res
}
