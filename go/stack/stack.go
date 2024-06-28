package main

// 本质上是一个反复写的slice，通过length控制栈顶
type Stack struct {
	Items    []string
	Length   int
	Capacity int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		Items:    make([]string, capacity, capacity),
		Length:   0,
		Capacity: capacity,
	}
}

// 入栈
func (s *Stack) Push(item string) bool {
	if s.Length >= s.Capacity {
		return false
	}
	s.Length++
	s.Items[s.Length-1] = item
	return true
}

// 出栈
func (s *Stack) Pop() string {
	if s.Length == 0 {
		return ""
	}
	item := s.Items[s.Length-1]
	s.Length--
	return item
}
