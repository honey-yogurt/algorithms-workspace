// see: https://leetcode.cn/problems/reverse-linked-list/
package linkedlist

// 1. 一遍迭代，每次遍历记录当前的 cur 以及 pre 以及 next，然后反转指针，重置 cur 以及 pre。  注意判断 cur != nil 比判断 cur.Next 好
// 2.

// O(n)
func reverseList1(head *SingleNode) *SingleNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre *SingleNode
	for cur.next != nil {
		// 保存下一个指针
		next := cur.next
		// 反转
		cur.next = pre
		// 当前 cur 变成 pre
		pre = cur
		// 下一个指针为循环的 cur
		cur = next
	}
	// 最后一个没有反转
	cur.next = pre
	return cur
}

func reverseList2(head *SingleNode) *SingleNode {
	cur := head
	var pre *SingleNode
	for cur != nil {
		// 保存下一个指针
		next := cur.next
		// 反转
		cur.next = pre
		// 当前 cur 变成 pre
		pre = cur
		// 下一个指针为循环的 cur
		cur = next
	}
	return pre
}
