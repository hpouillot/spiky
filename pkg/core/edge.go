package core

import (
	"math"
	"math/rand"
	"spiky/pkg/utils"
)

type Edge struct {
	weight float64
	delay  float64
	target *Neuron
	source *Neuron
}

func (edge *Edge) Forward(world *World) {
	world.Schedule(world.GetTime()+edge.delay, func(world *World) {
		target := edge.target
		target.potential += edge.weight
		if target.potential >= world.Const.Threshold {
			target.Fire(world)
		}
	})
}

func (edge *Edge) Backward(world *World) {
	preSpike := edge.source.GetLastSpikeTime()
	postSpike := edge.target.GetLastSpikeTime()
	deltaT := postSpike - preSpike
	Ap := 10.0
	Am := 5.0
	if deltaT >= 0 {
		edge.weight += Ap * world.Const.LearningRate * math.Exp(-deltaT/world.Const.Tho) * (world.Const.MaxWeight - edge.weight)
	} else {
		edge.weight -= Am * world.Const.LearningRate * math.Exp(deltaT/world.Const.Tho) * (edge.weight - world.Const.MinWeight)
	}
}

func NewEdge(source *Neuron, target *Neuron, cst *utils.Constants) *Edge {
	edge := &Edge{
		weight: rand.Float64() * cst.MaxWeight,
		delay:  rand.Float64() * cst.MaxDelay,
		source: source,
		target: target,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}
