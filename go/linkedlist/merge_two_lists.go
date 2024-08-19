// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// l1 和 l2 均按 非递减顺序 排列
package linkedlist


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个新的 list
	// 虚拟头结点
	dummy := &ListNode{
		-1,
		nil,
	}
	// 移动的指针，用来增添节点
	p := dummy
    // 双指针
	p1,p2 := list1,list2
	for p1 !=nil && p2 != nil {
		// 比较两个值大小
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next 
		}
		// p 移动
		p = p.Next
	}
	// 其中一个为nil
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	// 虚拟头结点不返回
	return dummy.Next
 }

