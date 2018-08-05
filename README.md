# gostu
Contains some implementations of simple data structures in Go.

## Stack
`gostu.Stack` represents a stack data structure.

```go
// initialize a new stack
s := NewStack()

// Push some values
s.Push(23)
s.Push("something else")
s.Push(34)

// Get top value
v, _ := s.Top()
// v => 34

// Pop value
v, _ := s.Pop()
// v => 34
// s.Top() => "something else"

// check if stack is emoty
s.IsEmpty()
// => false
```

Please note that you can add any kind of values to the stack. This is a limitation of golang, where there are no "type classes" supported for data structures. 