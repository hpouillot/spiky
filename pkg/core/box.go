package core

type Box[T interface{}] interface {
	// Get(point Point) T
	// Set(point Point, value T)
	// Add(value T) Point
	Visit(func(idx int, value *T))
	Size() int
	// Shape() []uint64
}

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
	neurons := make([]*Neuron, size)
	for i := 0; i < size; i++ {
		neurons[i] = NewNeuron(255.0, 0.1)
	}
	return &Layer{
		neurons: neurons,
	}
}
