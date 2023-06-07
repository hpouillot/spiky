package core

import (
	"math/rand"
	"spiky/pkg/utils"
)

type IEdge interface {
	Forward(world *World)
	Stdp(world *World, reward float64)
	GetId() string
	GetSource() *Neuron
	GetTarget() *Neuron
	Adjust(world *World, err float64)
}

type Edge struct {
	weight float64
	delay  float64
	target *Neuron
	source *Neuron
}

func (edge *Edge) Adjust(world *World, err float64) {
	spikeTime := edge.source.GetSpikeTime()
	if spikeTime != nil {
		deltaW := err * world.Const.LearningRate
		edge.AddWeight(deltaW, world)
	}
	edge.source.Adjust(world, err)
}

func (edge *Edge) Forward(world *World) {
	world.Schedule(world.GetTime()+edge.delay, func(world *World) {
		edge.target.Receive(world, edge.weight)
	})
}

func (edge *Edge) GetSource() *Neuron {
	return edge.source
}

func (edge *Edge) GetTarget() *Neuron {
	return edge.target
}

func (edge *Edge) GetId() string {
	return edge.source.id + edge.target.id
}

func (edge *Edge) AddWeight(deltaWeight float64, world *World) {
	edge.weight = utils.ClampFloat(edge.weight+deltaWeight, world.Const.MinWeight, world.Const.MaxWeight)
}

func (edge *Edge) Stdp(world *World, reward float64) {
	sourceTime := edge.source.GetSpikeTime()
	targetTime := edge.target.GetSpikeTime()
	if sourceTime != nil && targetTime != nil {
		var deltaT float64 = *targetTime - *sourceTime
		var deltaW float64 = 0.0
		if deltaT >= 0 {
			deltaW = reward
		}
		edge.AddWeight((deltaW * world.Const.LearningRate), world)
	}
}

type PositiveEdge struct {
	Edge
}

func (edge *PositiveEdge) SetWeight(newWeight float64, minWeight float64, maxWeight float64) {
	edge.weight = utils.ClampFloat(newWeight, 0, maxWeight)
}

func (edge *PositiveEdge) Forward(world *World) {
	world.Schedule(world.GetTime()+edge.delay, func(world *World) {
		edge.target.Receive(world, edge.weight)
	})
}

type NegativeEdge struct {
	Edge
}

func (edge *NegativeEdge) SetWeight(newWeight float64, minWeight float64, maxWeight float64) {
	edge.weight = utils.ClampFloat(newWeight, 0, -minWeight)
}

func (edge *NegativeEdge) Forward(world *World) {
	world.Schedule(world.GetTime()+edge.delay, func(world *World) {
		edge.target.Receive(world, -edge.weight)
	})
}

func NewEdge(source *Neuron, target *Neuron, cfg *ModelConfig) *Edge {
	edge := &Edge{
		delay:  rand.Float64() * cfg.MaxDelay,
		weight: 0,
		source: source,
		target: target,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}

func NewPositiveEdge(source *Neuron, target *Neuron, cfg *ModelConfig) *PositiveEdge {
	baseEdge := NewEdge(source, target, cfg)
	edge := &PositiveEdge{
		Edge: *baseEdge,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}

func NewNegativeEdge(source *Neuron, target *Neuron, cfg *ModelConfig) *NegativeEdge {
	baseEdge := NewEdge(source, target, cfg)
	edge := &NegativeEdge{
		Edge: *baseEdge,
	}
	source.synapses = append(source.synapses, edge)
	target.dendrites = append(target.dendrites, edge)
	return edge
}
