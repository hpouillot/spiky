package utils

import "testing"

func TestCallStack(t *testing.T) {
	stack := NewTimeStack[int]()
	stack.Push(0.0, 23)
	if stack.Len() != 1 {
		t.Errorf("Invalid stack length %v", stack.Len())
	}
	if stack.Pop() == nil {
		t.Error("Invalid item 1")
	}
	if stack.Pop() != nil {
		t.Error("Invalid item 2")
	}
}
