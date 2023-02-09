package main

func main() {

}

// https://leetcode.cn/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	dummy1 := new(ListNode)
	p1 := dummy1

	dummy2 := new(ListNode)
	p2 := dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		temp := head.Next
		head.Next = nil
		head = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
