package core

import (
	"spiky/pkg/utils"
)

type Neuron struct {
	potential         float64
	threshold         float64
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

func (node *Neuron) Fire(stack *utils.TimeStack) {
	node.spikes = append(node.spikes, stack.GetTime())
	for _, syn := range node.synapses {
		syn.Forward(stack)
	}
	// for _, dend := range node.dendrites {
	// 	stack.Push(stack.Delay(node.refrectory_period), dend.Backward)
	// }
	node.potential = 0
}

func NewNeuron(threshold float64, refractoryPeriod float64) *Neuron {
	return &Neuron{
		potential:         0.0,
		threshold:         threshold,
		refrectory_period: refractoryPeriod,
		spikes:            []float64{},
		synapses:          []*Edge{},
		dendrites:         []*Edge{},
	}
}
