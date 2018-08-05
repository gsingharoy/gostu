package gostu

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s.top != nil {
		t.Error("Expected top to be initialized to a null pointer")
	}
}
