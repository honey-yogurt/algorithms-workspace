package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {
	sll := NewSingleLinkedList()
	sll.Add(0, 10)
	sll.Add(0, 11)
	sll.Add(1, 12)
	sll.AddHeader(30)
	sll.AddHeader(17)
	sll.Add(4, 44)
	sll.AddTail(88)
	sll.Add(6, 98)
	sll.AddTail(100)
	sll.Print()                    // 17 30 11 12 44 10 98 88 100
	fmt.Println(sll.tail.val)      // 尾指针 100
	fmt.Println(sll.head.next.val) // 第一个实际的值 17
	sll.Delete(0)
	sll.Delete(7)
	sll.Delete(6)
	sll.Delete(5)
	sll.Delete(0)
	sll.Delete(3)
	sll.Delete(2)
	sll.Delete(1)
	sll.Print()                    // 11
	fmt.Println(sll.tail.val)      // 尾指针 11
	fmt.Println(sll.head.next.val) // 第一个实际的值 11
	sll.Delete(0)
	if sll.head == sll.tail {
		fmt.Println("ok") // ok
	}
}
