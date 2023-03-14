package utils

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

type BooleanQueue struct {
	Queue[bool]
}

func (q *BooleanQueue) Count() int {
	count := 0
	for _, i := range q.items {
		if i {
			count++
		}
	}
	return count
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
