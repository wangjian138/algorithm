package main

func main() {

}

// https://leetcode.cn/problems/3u1WK4/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB

	len1 := 0
	len2 := 0
	for p1 != nil {
		len1++
		p1 = p1.Next
	}

	for p2 != nil {
		len2++
		p2 = p2.Next
	}

	for len1 != len2 {
		if len1 > len2 {
			len1--
			headA = headA.Next
		} else {
			len2--
			headB = headB.Next
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
