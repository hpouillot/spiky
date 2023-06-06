package utils

import (
	"math"
)

type Queue[T interface{}] struct {
	size  int
	items []T
}

func (q *Queue[T]) Push(item T) {
	if len(q.items) > q.size {
		q.Pop()
	}
	q.items = append(q.items, item)
}

func (q *Queue[T]) Pop() {
	q.items = q.items[1:len(q.items)]
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) Visit(visitor func(item T)) {
	for _, i := range q.items {
		visitor(i)
	}
}

type BooleanQueue struct {
	Queue[bool]
}

func (q *BooleanQueue) Count() int {
	count := 0
	q.Visit(func(item bool) {
		if item {
			count++
		}
	})
	return count
}

type FloatQueue struct {
	Queue[float64]
}

func (q *FloatQueue) Mean() float64 {
	sum := 0.0
	q.Visit(func(item float64) {
		sum += item
	})
	return sum / float64(q.Len())
}

func (q *FloatQueue) Var(mean float64) float64 {
	sum := 0.0
	q.Visit(func(item float64) {
		diff := item - mean
		sum += diff * diff
	})
	return sum / float64(q.Len())
}

func (q *FloatQueue) Scale(value float64) float64 {
	if q.Len() == 0 {
		return 1.0
	}
	mean := q.Mean()
	variance := q.Var(mean)
	if variance <= 0.00000001 {
		return 1.0
	}
	std := math.Sqrt(variance)
	return (value - mean) / std
}

func NewQueue[T interface{}](maxSize int) *Queue[T] {
	return &Queue[T]{
		size:  maxSize,
		items: []T{},
	}
}

func NewBooleanQueue(maxSize int) *BooleanQueue {
	return &BooleanQueue{
		Queue: *NewQueue[bool](maxSize),
	}
}

func NewFloatQueue(maxSize int) *FloatQueue {
	return &FloatQueue{
		Queue: *NewQueue[float64](maxSize),
	}
}
