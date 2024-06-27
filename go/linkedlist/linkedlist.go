package linkedlist

import (
	"fmt"
)

// 1. 同时持有头尾结点的引用
// 2. 虚拟头尾结点(双链表)
// 3. 处理内存泄漏
// 4. pos 是索引，从 0 开始，不包括虚拟head

// 单链表结点
type SingleNode struct {
	next *SingleNode
	val  any
}

type SingleLinkedList struct {
	head *SingleNode // 虚拟头结点
	tail *SingleNode // 实际的尾部结点引用
	size int         // 实际结点，不计虚拟头结点
}

func NewSingleLinkedList() *SingleLinkedList {
	head := &SingleNode{nil, nil}
	return &SingleLinkedList{
		head: head,
		tail: head,
		size: 0,
	}
}

// 头插法
func (s *SingleLinkedList) AddHeader(v any) {
	n := &SingleNode{
		next: s.head.next, // 头插法，新建的节点指向 head 的 next
		val:  v,
	}
	s.head.next = n
	// 只有头插法在插入第一个结点时候要变动 tail
	// 因为一旦 tail 有实际的node了，怎么头插，tail都不会改变的，所以不用改变
	if s.size == 0 {
		s.tail = n
	}
	s.size++
}

// 尾插法
func (s *SingleLinkedList) AddTail(v any) {
	n := &SingleNode{
		next: nil,
		val:  v,
	}
	s.tail.next = n
	s.tail = n
	s.size++
}

func (s *SingleLinkedList) Add(pos int, v any) bool {
	// 查找上一个结点
	var tar *SingleNode
	if pos == 0 {
		tar = s.head
	} else {
		tar = s.GetNode(pos - 1)
		if tar == nil {
			return false
		}
	}
	n := &SingleNode{
		next: tar.next, // 指向 pos
		val:  v,
	}
	tar.next = n // 指向当前
	if s.size == 0 {
		s.tail = n // 注意调整尾指针
	}
	s.size++
	return true
}

func (s *SingleLinkedList) Delete(pos int) bool {
	if s.size == 0 {
		return false
	}
	// 查找上一个结点
	var tar *SingleNode
	if pos == 0 {
		tar = s.head
	} else {
		tar = s.GetNode(pos - 1)
		if tar == nil {
			return false
		}
	}
	// n := tar.next
	tar.next = tar.next.next
	if pos == s.size-1 { // 只有删除最后一个结点时候需要调整尾指针
		s.tail = tar
	}
	s.size--
	// n = nil  回收内存，但是go变量必须使用，所以注释
	return true
}

func (s *SingleLinkedList) Print() {
	cur := s.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.val)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (s *SingleLinkedList) GetNode(pos int) *SingleNode {
	if !s.checkPos(pos) {
		return nil
	}
	cur := s.head
	for i := 0; i <= pos; i++ {
		cur = cur.next
	}
	return cur
}

func (s *SingleLinkedList) checkPos(pos int) bool {
	if pos < 0 || pos > s.size {
		return false
	}
	return true
}
