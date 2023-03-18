package core

import (
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
		edge.target.Receive(world)
	})
}

func (edge *Edge) Adjust(world *World, err float64) {
	_, preErr := edge.source.GetLastSpikeTime()
	if preErr != nil {
		// source did not spike
		return
	}
	deltaW := err * world.Const.LearningRate
	edge.weight = utils.ClampFloat(edge.weight+deltaW, world.Const.MinWeight, world.Const.MaxWeight)
}

func NewEdge(source *Neuron, target *Neuron, cst *utils.Constants) *Edge {
	edge := &Edge{
		delay:  rand.Float64() * cst.MaxDelay,
		weight: utils.ClampFloat(rand.NormFloat64()*(cst.MaxWeight/2), cst.MinWeight, cst.MaxWeight),
		source: source,
		target: target,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}
