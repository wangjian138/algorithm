package main

func main() {

}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/submissions/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head

	for n > 0 {
		p = p.Next
		n--
	}

	dummp := &ListNode{Next: head}
	p1 := dummp
	for p != nil {
		p1 = p1.Next
		p = p.Next
	}
	p1.Next = p1.Next.Next

	return dummp.Next
}
