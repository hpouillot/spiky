package core

import "github.com/jaevor/go-nanoid"

type Layer struct {
	name    string
	neurons []*Neuron
}

func (nb *Layer) Visit(fn func(idx int, value *Neuron)) {
	for idx, n := range nb.neurons {
		fn(idx, n)
	}
}

func (nb *Layer) Get(idx int) *Neuron {
	return nb.neurons[idx]
}

func (nb *Layer) Size() int {
	return len(nb.neurons)
}

func (nb *Layer) GetName() string {
	return nb.name
}

func NewLayer(name string, size int) *Layer {
	idGenerator, err := nanoid.Standard(21)
	if err != nil {
		panic("Can't instantiate nanoid")
	}
	neurons := make([]*Neuron, size)
	for i := 0; i < size; i++ {
		neurons[i] = NewNeuron(idGenerator())
	}
	return &Layer{
		name:    name,
		neurons: neurons,
	}
}
