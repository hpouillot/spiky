package core

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Neuron struct {
	id        string
	potential float64
	spikes    []float64
	synapses  []*Edge
	dendrites []*Edge
}

func (node *Neuron) GetSpikes() *[]float64 {
	return &node.spikes
}

func (node *Neuron) GetLastSpikeTime() (float64, error) {
	spikeLength := len(node.spikes)
	if spikeLength == 0 {
		return 0, errors.New("no spike")
	}
	return node.spikes[len(node.spikes)-1], nil
}

func (node *Neuron) Fire(world *World) {
	node.spikes = append(node.spikes, world.GetTime())
	logrus.Info("node fired")
	for _, syn := range node.synapses {
		syn.Forward(world)
	}
	// for _, dend := range node.dendrites {
	// 	world.Schedule(world.GetTime()+2.0, dend.Backward)
	// }
	node.potential = 0
}

func (node *Neuron) Adjust(world *World, err float64) {
	for _, dend := range node.dendrites {
		dend.Adjust(world, err)
	}
}

func (n *Neuron) Clear() {
	n.potential = 0
	n.spikes = []float64{}
}

func NewNeuron(id string) *Neuron {
	return &Neuron{
		id:        id,
		potential: 0.0,
		spikes:    []float64{},
		synapses:  []*Edge{},
		dendrites: []*Edge{},
	}
}
