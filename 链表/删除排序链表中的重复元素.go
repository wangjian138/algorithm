package main

func main() {

}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil {
		if fast.Val == slow.Val {
			fast = fast.Next
			// 最后一个
			if fast == nil {
				slow.Next = fast
			}
			continue
		}
		slow.Next = fast
		slow = fast
		fast = fast.Next
	}
	return head
}
