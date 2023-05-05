package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	list := make([]*Node, 0)
	list = append(list, root)
	for len(list) > 0 {
		temp := make([]*Node, 0)
		listLen := len(list)
		for i := 0; i < listLen; i++ {
			node := list[i]
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
			if i+1 == listLen {
				node.Next = nil
				continue
			}
			node.Next = list[i+1]
		}
		list = temp
	}

	return root
}
