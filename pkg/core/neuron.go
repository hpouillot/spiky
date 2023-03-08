package core

import (
	"github.com/sirupsen/logrus"
)

type Neuron struct {
	potential         float64
	refrectory_period float64
	spikes            []float64
	synapses          []*Edge
	dendrites         []*Edge
}

func (node *Neuron) GetSpikes() *[]float64 {
	return &node.spikes
}

func (node *Neuron) GetLastSpikeTime() float64 {
	return node.spikes[len(node.spikes)-1]
}

func (node *Neuron) Fire(world *World) {
	node.spikes = append(node.spikes, world.GetTime())
	logrus.Info("node fired")
	for _, syn := range node.synapses {
		syn.Forward(world)
	}
	for _, dend := range node.dendrites {
		world.Schedule(world.GetTime()+world.Const.RefractoryPeriod, dend.Backward)
	}
	node.potential = 0
}

func NewNeuron() *Neuron {
	return &Neuron{
		potential: 0.0,
		spikes:    []float64{},
		synapses:  []*Edge{},
		dendrites: []*Edge{},
	}
}
