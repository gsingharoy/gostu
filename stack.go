package gostu

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
