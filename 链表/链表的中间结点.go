package main

func main() {

}

// https://leetcode.cn/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	p := head

	for head != nil {
		head = head.Next
		if head == nil {
			return p
		}
		head = head.Next
		p = p.Next
	}

	return p
}
