// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。
package linkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func partition(head *ListNode, x int) *ListNode {
	// 存放小于 x 的虚拟节点
	dummy1 := &ListNode{
		-1,
		nil,
	}

	// 存放大于等于 x 的虚拟节点
	dummy2 := &ListNode{
		-1,
		nil,
	}

	p := head
	p1 := dummy1
	p2 := dummy2

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
			
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 不能直接让 p 指针前进，
		// p = p.Next
		// 断开原链表中的每个节点的 next 指针
		temp := p.Next
		p.Next = nil
		p = temp
	}

	// mege 两个链表
	p1.Next = dummy2.Next

	return dummy1.Next
 }



 func partition2(head *ListNode, x int) *ListNode {
	// 存放小于 x 的虚拟节点
	dummy1 := &ListNode{
		-1,
		nil,
	}

	// 存放大于等于 x 的虚拟节点
	dummy2 := &ListNode{
		-1,
		nil,
	}

	p := head
	p1 := dummy1
	p2 := dummy2

	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
			
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		p = p.Next
	}

	// mege 两个链表
	p1.Next = dummy2.Next
	// 考虑环出现的情况
	p2.Next = nil

	return dummy1.Next
}