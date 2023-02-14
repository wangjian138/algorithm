package main

func main() {

}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	nxt := head
	k1 := k
	for k1 > 0 {
		k1--
		if nxt == nil {
			return head
		}
		nxt = nxt.Next
	}

	start := reverse(cur, nxt)
	cur.Next = reverseKGroup(nxt, k)
	return start
}

func reverse(a, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	cur = a
	nxt = a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
