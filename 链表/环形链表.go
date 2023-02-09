package main

func main() {

}

// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	p := head
	for head != nil && head.Next != nil {
		head = head.Next.Next
		p = p.Next
		if p == head {
			return true
		}
	}

	return false
}
