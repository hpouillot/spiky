package utils

import (
	"container/heap"
)

type StackItem[T interface{}] struct {
	Value T
	Time  float64
	index int
}

type TimeStack[T interface{}] struct {
	heap itemHeap[T]
}

func NewTimeStack[T interface{}]() *TimeStack[T] {
	return &TimeStack[T]{
		heap: itemHeap[T]{},
	}
}

func (s *TimeStack[T]) Push(time float64, process T) {
	newItem := &StackItem[T]{
		Value: process,
		Time:  time,
	}
	heap.Push(&s.heap, newItem)
}

func (s *TimeStack[T]) Pop() *StackItem[T] {
	if s.heap.Len() == 0 {
		return nil
	}
	return heap.Pop(&s.heap).(*StackItem[T])
}

func (s *TimeStack[T]) Len() int {
	return s.heap.Len()
}

func (s *TimeStack[T]) Reset() {
	s.heap = nil
}

type itemHeap[T interface{}] []*StackItem[T]

func (ih *itemHeap[T]) Len() int {
	return len(*ih)
}

func (ih *itemHeap[T]) Less(i, j int) bool {
	return (*ih)[i].Time < (*ih)[j].Time
}

func (ih *itemHeap[T]) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap[T]) Push(x interface{}) {
	it := x.(*StackItem[T])
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap[T]) Pop() interface{} {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
