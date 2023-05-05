package main

func main() {

}

// https://leetcode.cn/problems/lru-cache/
type LRUCache struct {
	cap        int
	nodeMap    map[int]*node
	doubleNode *DoubleList
}

type node struct {
	pre      *node
	next     *node
	val, key int
}

// https://leetcode-cn.com/problems/lru-cache/
func Constructor(capacity int) LRUCache {
	c := LRUCache{}

	c.cap = capacity

	nodeMap := make(map[int]*node, capacity)

	c.nodeMap = nodeMap

	dN := new(DoubleList)

	tail := new(node)
	head := new(node)

	head.next = tail
	tail.pre = head

	dN.head = head
	dN.tail = tail

	c.doubleNode = dN

	return c
}

func (this *LRUCache) Get(key int) int {
	node, res := this.nodeMap[key]
	if res == false {
		return -1
	}

	this.doubleNode.makeRecently(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	//判断老的key是否存在
	olNode, res := this.nodeMap[key]
	if res == true {
		olNode.val = value
		this.doubleNode.makeRecently(olNode)
		return
	}

	//新的node
	newNode := &node{
		key: key,
		val: value,
	}

	this.nodeMap[key] = newNode

	//判断是否已经达到上限
	if this.cap <= this.doubleNode.size {
		//移除掉最老的
		rb := this.doubleNode.removeFirst()
		//删除掉map
		delete(this.nodeMap, rb.key)
	}

	//加入新的
	this.doubleNode.addLast(newNode)

	return
}

type DoubleList struct {
	head, tail *node
	size       int
}

func (this *DoubleList) makeRecently(node *node) {
	this.remove(node)
	this.addLast(node)
}

func (this *DoubleList) remove(node *node) {
	node.pre.next = node.next
	node.next.pre = node.pre

	this.size--
}

func (this *DoubleList) addLast(node *node) {

	node.pre = this.tail.pre
	node.next = this.tail

	this.tail.pre.next = node

	this.tail.pre = node

	this.size++
}

func (this *DoubleList) removeFirst() *node {
	if this.head.next == this.tail {
		return nil
	}

	first := this.head.next

	this.remove(first)

	return first
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
