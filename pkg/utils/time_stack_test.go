package utils

import "testing"

func TestCallStack(t *testing.T) {
	items := []int32{}
	stack := NewTimeStack()
	stack.Push(0.0, Process(func(s *TimeStack) {
		items = append(items, 20)
		s.Push(1.1, Process(func(stack *TimeStack) {
			items = append(items, 30)
		}))
	}))
	stack.Resolve(0, 3)
	if len(items) != 2 {
		t.Error("Invalid items size {}", len(items))
	}
	if items[0] != 20 || items[1] != 30 {
		t.Error("Invalid elements")
	}
}
