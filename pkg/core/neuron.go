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
	return neuron.spikes[len((neuron.spikes))-1], nil
}

func (neuron *Neuron) Fire(world *World) {
	neuron.spikes = append(neuron.spikes, world.GetTime())
	world.markDirty(neuron)
	for _, syn := range neuron.synapses {
		syn.Forward(world)
	}
	neuron.potential = 0
}

func (neuron *Neuron) Receive(world *World) {
	_, err := neuron.GetLastSpikeTime()
	if err == nil {
		return
	}
	potential := neuron.getPotential(world)
	if potential >= world.Const.Threshold {
		neuron.Fire(world)
	}
}

func (neuron *Neuron) Adjust(world *World, err float64) {
	for _, dend := range neuron.dendrites {
		dend.Adjust(world, err)
	}
}

func (neuron *Neuron) getPotential(world *World) float64 {
	potential := 0.0
	currenTime := world.GetTime()
	for _, dend := range neuron.dendrites {
		lastSpikeTime, err := dend.source.GetLastSpikeTime()
		if err == nil {
			potential += (1 - (lastSpikeTime-currenTime)/world.Const.MaxTime) * dend.weight
		}
	}
	return potential
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
