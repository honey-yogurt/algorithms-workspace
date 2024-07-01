package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack(5)
	s.Push("hello")
	s.Push("world")
	s.Pop()
	s.Push("go")
	s.Push("not")
	s.Push("java")
	s.Push("rust")
	s.Push("c")          // false
	fmt.Println(s.Pop()) // rust
}
