package main

func main() {

}

// https://leetcode.cn/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successNode *ListNode

func reverseN(head *ListNode, num int) *ListNode {
	if num == 1 {
		successNode = head.Next
		return successNode
	}
	last := reverseN(head.Next, num-1)
	head.Next.Next = head
	head.Next = successNode
	return last
}
