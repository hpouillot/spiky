package core

type Node interface {
	GetId() string
	Compute(time Time, queue *Queue)
	SetSpike(time Time, spiked bool)
	GetSpike(time Time) bool
	GetPosition() Point
	GetSpikeRate(startTime Time, endTime Time) (float64, error)
	Connect(node Node) Edge
	AddDendrite(edge Edge)
	AddSynapse(edge Edge)
	GetSynapses() []Edge
	GetDendrites() []Edge
	GetParents() []Node
	GetChildren() []Node
	GetLastSpikeTime() Time
	GetSpikeTimes(startTime Time, endTime Time) []Time
	Reset()
}
