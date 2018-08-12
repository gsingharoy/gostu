package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := New()
	if s.top != nil {
		t.Error("Expected top to be initialized to a null pointer")
	}
}

func TestStackPush(t *testing.T) {
	t.Log("adding to an empty stack")
	s := New()
	err1 := s.Push(45)
	if err1 != nil {
		t.Error("expected no error")
	}
	if s.top.value != 45 {
		t.Error("expected value to be added")
	}
	if s.top.next != nil {
		t.Error("expected next pointer to be nil")
	}

	t.Log("adding to an existing Stack")
	err2 := s.Push(56)
	if err2 != nil {
		t.Error("expected no error")
	}
	if s.top.value != 56 {
		t.Error("expected value to be added")
	}
	if s.top.next == nil {
		t.Error("expected next pointer")
	}
	if s.top.next.value != 45 && s.top.next.next != nil {
		t.Error("expected correct next value")
	}
}

func TestStackPop(t *testing.T) {
	t.Log("popping from an empty stack")
	s := New()
	_, err := s.Pop()
	if err != ErrEmptyStack {
		t.Error("expected error")
	}
	s.Push(45)
	s.Push(56)
	s.Push(69)
	v1, _ := s.Pop()
	if v1 != 69 {
		t.Error("expected last value to be popped")
	}

	v2, _ := s.Pop()
	if v2 != 56 {
		t.Error("expected last value to be popped")
	}

	v3, _ := s.Pop()
	if v3 != 45 {
		t.Error("expected last value to be popped")
	}
}

func TestStackIsEmpty(t *testing.T) {
	t.Log("when stack is empty")

	s := New()
	if !s.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	s.Push(34)

	if s.IsEmpty() {
		t.Error("expected stack to be not empty")
	}
}

func TestStackTop(t *testing.T) {
	t.Log("when stack is empty")

	s := New()

	_, err := s.Top()
	if err != ErrEmptyStack {
		t.Error("expected ErrEmptyStack to be returned")
	}
}
