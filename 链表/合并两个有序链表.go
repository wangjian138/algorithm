package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// https://leetcode.cn/problems/merge-two-sorted-lists/submissions/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := new(ListNode)
	p1 := p
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p1.Next = list1
			list1 = list1.Next
		} else {
			p1.Next = list2
			list2 = list2.Next
		}
		p1 = p1.Next
	}
	if list1 != nil {
		p1.Next = list1
	}
	if list2 != nil {
		p1.Next = list2
	}

	return p.Next
}
