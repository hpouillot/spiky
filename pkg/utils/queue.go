package utils

import (
	"spiky/pkg/core"

	"github.com/aidarkhanov/nanoid/v2"
	"github.com/wangjia184/sortedset"
)

type Queue struct {
	orderedSet *sortedset.SortedSet
}

func (q *Queue) Add(time core.Time, node core.Node) {
	str, _ := nanoid.New()
	q.orderedSet.AddOrUpdate(str, sortedset.SCORE(time), node)
}

func (q *Queue) Count() int {
	return q.orderedSet.GetCount()
}

func (q *Queue) Pop() (core.Time, core.Node) {
	item := q.orderedSet.PopMin()
	return core.Time(item.Score()), item.Value.(core.Node)
}

func NewQueue() *Queue {
	return &Queue{
		orderedSet: sortedset.New(),
	}
}
