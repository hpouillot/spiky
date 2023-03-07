package core

import (
	"math"
	"spiky/pkg/utils"
)

type Edge struct {
	weight       float64
	delay        float64
	target       *Neuron
	source       *Neuron
	tho          float64
	maxWeight    float64
	minWeight    float64
	learningRate float64
}

func (edge *Edge) Forward(stack *utils.TimeStack) {
	stack.Push(stack.Delay(edge.delay), func(stack *utils.TimeStack) {
		target := edge.target
		target.potential += edge.weight
		if target.potential >= target.threshold {
			target.Fire(stack)
		}
	})
}

func (edge *Edge) Backward(stack *utils.TimeStack) {
	preSpike := edge.source.GetLastSpikeTime()
	postSpike := edge.target.GetLastSpikeTime()
	deltaT := postSpike - preSpike
	Ap := 10.0
	Am := 5.0
	if deltaT >= 0 {
		edge.weight += Ap * edge.learningRate * math.Exp(-deltaT/edge.tho) * (edge.maxWeight - edge.weight)
	} else {
		edge.weight -= Am * edge.learningRate * math.Exp(deltaT/edge.tho) * (edge.weight - edge.minWeight)
	}
}

func NewEdge(source *Neuron, target *Neuron) *Edge {
	edge := &Edge{
		weight:       1000.0,
		delay:        0.1,
		source:       source,
		target:       target,
		tho:          10,
		maxWeight:    250.0,
		minWeight:    0.0,
		learningRate: 0,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}
