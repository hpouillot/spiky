package core

import "github.com/jaevor/go-nanoid"

type Layer struct {
	neurons []*Neuron
}

func (nb *Layer) Visit(fn func(idx int, value *Neuron)) {
	for idx, n := range nb.neurons {
		fn(idx, n)
	}
}

func (nb *Layer) Size() int {
	return len(nb.neurons)
}

func NewLayer(size int) *Layer {
	idGenerator, err := nanoid.Standard(21)
	if err != nil {
		panic("Can't instantiate nanoid")
	}
	neurons := make([]*Neuron, size)
	for i := 0; i < size; i++ {
		neurons[i] = NewNeuron(idGenerator())
	}
	return &Layer{
		neurons: neurons,
	}
}
