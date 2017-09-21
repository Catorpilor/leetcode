package utils

// Stack is LIFO
type Stack struct {
	top  *element
	size int
}

type element struct {
	value interface{}
	next  *element
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Len returns a stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push pushes a value to the stack
func (s *Stack) Push(value interface{}) {
	s.top = &element{value: value, next: s.top}
	s.size += 1
}

// Top returns the top element in the stack
func (s *Stack) Top() interface{} {
	return s.top.value
}

// Pop pops the top element in the stack
func (s *Stack) Pop() interface{} {
	var value interface{}
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size -= 1
	}
	return value
}

// NewStack returns a brand new stack
func NewStack() *Stack {
	return &Stack{}
}
