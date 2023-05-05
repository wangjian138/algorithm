package main

func main() {

}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 注意：go 代码由 chatGPT🤖 根据我的 java 代码翻译，旨在帮助不同背景的读者理解算法逻辑。
// 本代码已经通过力扣的测试用例，应该可直接成功提交。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// p1 指向 A 链表头结点，p2 指向 B 链表头结点
	p1, p2 := headA, headB
	for p1 != p2 {
		// p1 走一步，如果走到 A 链表末尾，转到 B 链表
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// 详细解析参见：
// https://labuladong.github.io/article/?qno=160
