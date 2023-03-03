package models

import (
	"spiky/pkg/core"
	"spiky/pkg/utils"
)

type baseModel struct {
	input  core.Layer
	output core.Layer
}

// Single sample run
func (m *baseModel) Run(duration int) {
	queue := utils.NewQueue()
	time := core.Time(0)
	end_time := core.Time(duration)
	m.input.Visit(func(node core.Node, _ int) {
		queue.Add(time, node)
	})
	for queue.Count() != 0 && time < end_time {
		newTime, newNode := queue.Pop()
		time = newTime
		newNode.Compute(time, queue)
	}
}

func (m *baseModel) Train(duration int) {

}

func (m *baseModel) Reset() {
	m.Visit(func(node core.Node, idx int) {
		node.Reset()
	})
}

func (m *baseModel) Visit(fn func(node core.Node, idx int)) {
	visitedNodes := make(map[string]bool)
	nodesToVisit := make([]core.Node, 0)
	m.input.Visit(func(node core.Node, idx int) {
		nodesToVisit = append(nodesToVisit, node)
	})
	idx := 0
	for len(nodesToVisit) != 0 {
		currentNode := nodesToVisit[0]
		nodesToVisit = nodesToVisit[1:]
		visitedNodes[currentNode.GetId()] = true
		fn(currentNode, idx)
		for _, child := range currentNode.GetChildren() {
			if !visitedNodes[child.GetId()] {
				nodesToVisit = append(nodesToVisit, child)
			}
		}
		idx++
	}
}

func Model(input core.Layer, output core.Layer) core.Model {
	return &baseModel{
		input:  input,
		output: output,
	}
}
