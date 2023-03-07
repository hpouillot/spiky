package utils

import (
	"container/heap"
	"fmt"
)

type Process (func(stack *TimeStack))
type Time (float64)

type TimeStack struct {
	heap *itemHeap
	time float64
}

func NewTimeStack() *TimeStack {
	return &TimeStack{
		heap: &itemHeap{},
		time: 0.0,
	}
}

func (s *TimeStack) Delay(delay float64) float64 {
	return s.time + delay
}

func (s *TimeStack) GetTime() float64 {
	return s.time
}

func (s *TimeStack) setTime(time float64) {
	s.time = time
}

func (s *TimeStack) Push(time float64, process Process) {
	newItem := &item{
		value: process,
		time:  time,
	}
	heap.Push(s.heap, newItem)
}

func (s *TimeStack) Resolve(start float64, end float64) error {
	if s.heap.Len() == 0 {
		return nil
	}
	s.setTime(start)
	for s.heap.Len() != 0 && s.GetTime() < end {
		stackItem := heap.Pop(s.heap).(*item)
		stackItem.value(s)
		s.setTime(stackItem.time)
		fmt.Println("Current Time", s.GetTime())
	}
	return nil
}

func (s *TimeStack) Clear() {
	s.heap = &itemHeap{}
}

type itemHeap []*item

type item struct {
	value Process
	time  float64
	index int
}

func (ih *itemHeap) Len() int {
	return len(*ih)
}

func (ih *itemHeap) Less(i, j int) bool {
	return (*ih)[i].time < (*ih)[j].time
}

func (ih *itemHeap) Swap(i, j int) {
	(*ih)[i], (*ih)[j] = (*ih)[j], (*ih)[i]
	(*ih)[i].index = i
	(*ih)[j].index = j
}

func (ih *itemHeap) Push(x interface{}) {
	it := x.(*item)
	it.index = len(*ih)
	*ih = append(*ih, it)
}

func (ih *itemHeap) Pop() interface{} {
	old := *ih
	item := old[len(old)-1]
	*ih = old[0 : len(old)-1]
	return item
}
