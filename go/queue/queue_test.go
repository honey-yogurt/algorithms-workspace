package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	aq := NewArrayQueue(5)
	aq.Enqueue("this")
	aq.Enqueue("is")
	aq.Enqueue("a")
	aq.Enqueue("book")
	aq.Enqueue("!")
	fmt.Println(aq.Enqueue("?")) // false , 不起作用
	fmt.Println(aq.Dequeue())    // this true
	aq.Enqueue("this")
	fmt.Println(aq.Dequeue()) // is true
	aq.Dequeue()
	aq.Dequeue()
	aq.Dequeue()
	fmt.Println(aq.Dequeue()) // this true
	fmt.Println(aq.Dequeue()) // "" false
	aq.Enqueue("we are new")
	fmt.Println(aq.Dequeue()) // we are new true
}
