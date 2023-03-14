package core

import (
	"errors"
)

type Neuron struct {
	id        string
	potential float64
	spikes    []float64
	synapses  []*Edge
	dendrites []*Edge
}

func (neuron *Neuron) GetSpikes() []float64 {
	return neuron.spikes
}

func (neuron *Neuron) GetLastSpikeTime() (float64, error) {
	spikeLength := len(neuron.spikes)
	if spikeLength == 0 {
		return 0, errors.New("no spike")
	}
	return (neuron.spikes)[len((neuron.spikes))-1], nil
}

func (neuron *Neuron) Fire(world *World) {
	neuron.spikes = append(neuron.spikes, world.GetTime())
	world.markDirty(neuron)
	for _, syn := range neuron.synapses {
		syn.Forward(world)
	}
	neuron.potential = 0
}

func (neuron *Neuron) Adjust(world *World, err float64) {
	for _, dend := range neuron.dendrites {
		dend.Adjust(world, err)
	}
}

func (n *Neuron) Reset() {
	n.potential = 0
	n.spikes = nil
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
