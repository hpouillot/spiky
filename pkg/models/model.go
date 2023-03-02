package models

import (
	"spiky/pkg/core"
)

type baseModel struct {
	input  core.Layer
	output core.Layer
}

// Single sample run
func (m *baseModel) Run(duration int) {
	queue := core.NewQueue()
	time := core.Time(0)
	end_time := core.Time(duration)
	m.input.Visit(func(node core.Node) {
		queue.Add(time, node)
	})
	for queue.GetCount() != 0 && time < end_time {
		newTime, newNode := queue.PopMin()
		time = newTime
		newNode.Compute(time, queue)
	}
}

func (m *baseModel) Train(duration int) {

}

func Model(input core.Layer, output core.Layer) core.Model {
	return &baseModel{
		input:  input,
		output: output,
	}
}
