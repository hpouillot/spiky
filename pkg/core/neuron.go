package core

type Neuron struct {
	id        string
	potential float64
	spikeTime *float64
	synapses  []IEdge
	dendrites []IEdge
}

func (neuron *Neuron) GetSpikeTime() *float64 {
	return neuron.spikeTime
}

func (n *Neuron) SetSpikeTime(world *World, time *float64) {
	n.spikeTime = time
	world.markDirty(n)
}

func (n *Neuron) Potentiate(world *World, delta float64) {
	n.potential = n.potential + delta
	world.markDirty(n)
}

func (neuron *Neuron) Fire(world *World) {
	spikeTime := world.GetTime()
	neuron.SetSpikeTime(world, &spikeTime)
	for _, syn := range neuron.synapses {
		syn.Forward(world)
	}
	neuron.potential = 0
}

func (neuron *Neuron) Receive(world *World, signal float64) {
	if neuron.spikeTime != nil {
		return
	}
	neuron.Potentiate(world, signal)
	if neuron.potential >= world.Const.Threshold {
		neuron.Fire(world)
	}
}

func (neuron *Neuron) Adjust(world *World, err float64) {
	for _, dend := range neuron.dendrites {
		dend.Adjust(world, err)
	}
}

func (n *Neuron) Reset() {
	n.potential = 0
	n.spikeTime = nil
}

func NewNeuron(id string) *Neuron {
	return &Neuron{
		id:        id,
		potential: 0.0,
		spikeTime: nil,
		synapses:  []IEdge{},
		dendrites: []IEdge{},
	}
}
