package main

func main() {

}

// https://leetcode.cn/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	reverseNode := reverse(slow)
	for head != nil && reverseNode != nil {
		if head.Val != reverseNode.Val {
			return false
		}
		head = head.Next
		reverseNode = reverseNode.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
