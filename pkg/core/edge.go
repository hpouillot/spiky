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
		edge.target.Receive(world, edge.weight)
	})
}

func (edge *Edge) Adjust(world *World, err float64) {
	spikeTime := edge.source.GetSpikeTime()
	if spikeTime != nil {
		deltaW := err * world.Const.LearningRate
		edge.weight = utils.ClampFloat(edge.weight+deltaW, world.Const.MinWeight, world.Const.MaxWeight)
	}
	edge.source.Adjust(world, err)
}

func (edge *Edge) Stdp(world *World, reward float64) {
	sourceTime := edge.source.GetSpikeTime()
	targetTime := edge.target.GetSpikeTime()
	if sourceTime != nil && targetTime != nil {
		var deltaT float64 = *targetTime - (*sourceTime + edge.delay)
		var deltaW float64 = 0.0
		if deltaT >= 0 {
			deltaW = reward
		} else {
			deltaW = -reward
		}
		edge.weight = utils.ClampFloat(edge.weight+(deltaW*world.Const.LearningRate), world.Const.MinWeight, world.Const.MaxWeight)
	}
}

func NewEdge(source *Neuron, target *Neuron, cfg *ModelConfig) *Edge {
	edge := &Edge{
		delay:  rand.Float64() * cfg.MaxDelay,
		weight: rand.Float64() * cfg.MaxWeight,
		source: source,
		target: target,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}
