package gostu

// Simple Stack data structure
// This is a unidirectional simple data structure
// read here to know more about stacks : https://www.geeksforgeeks.org/stack-data-structure/
type Stack struct {
	top *stackField
}

type stackField struct {
	value interface{}
	next  *stackField
}

// NewStack() returns a pointer to a newly created empty Stack instance
func NewStack() *Stack {
	return &Stack{}
}

// Push(v interface{}) returns an error value
// This method 'pushed' a value to the stack
func (s *Stack) Push(v interface{}) error {
	newField := &stackField{value: v, next: s.top}
	s.top = newField
	return nil
}

// Pop() returns the top value and an error if it exists
// This method then sets the top to the next value
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}

// IsEmpty() returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Top() returns the top value of the Stack
// Returns an error if the top is empty
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	return s.top.value, nil
}
