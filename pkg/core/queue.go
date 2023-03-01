package core

import "github.com/wangjia184/sortedset"

type Queue struct {
	orderedSet *sortedset.SortedSet
}

func (q *Queue) Add(time Time, node Node) {
	q.orderedSet.AddOrUpdate(node.GetId(), sortedset.SCORE(time), node)
}

func (q *Queue) GetCount() int {
	return q.orderedSet.GetCount()
}

func (q *Queue) PopMin() (Time, Node) {
	item := q.orderedSet.PopMin()
	return Time(item.Score()), item.Value.(Node)
}

func NewQueue() *Queue {
	return &Queue{
		orderedSet: sortedset.New(),
	}
}
