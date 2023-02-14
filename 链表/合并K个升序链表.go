package main

func main() {
	mergeKLists([]*ListNode{
		&ListNode{Val: 1},
		&ListNode{Val: 1},
	})
}

// https://leetcode.cn/problems/merge-k-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeDepth(lists, 0, len(lists)-1)
}

func mergeDepth(lists []*ListNode, l, r int) *ListNode {
	if l == r-1 {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return merge(mergeDepth(lists, l, mid), mergeDepth(lists, mid+1, r))
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
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
