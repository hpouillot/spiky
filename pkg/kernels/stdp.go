package kernels

import (
	"math"
	"spiky/pkg/core"
)

type StdpKernel struct {
	Threshold float64
	Tho       float64
}

func (m *StdpKernel) GetContrib(edge core.Edge, startTime core.Time, endTime core.Time) float64 {
	source := edge.GetSource()
	contrib := 0.0

	for _, time := range source.GetSpikeTimes(startTime, endTime) {
		contrib += edge.GetWeight() * math.Exp(-(float64(endTime-time))/m.Tho)
	}

	return contrib
}

func (m *StdpKernel) Compute(node core.Node, time core.Time, queue *core.Queue) {
	potential := 0.0
	spiked := false
	startTime := node.GetLastSpikeTime()
	for _, dendrite := range node.GetDendrites() {
		potential += m.GetContrib(dendrite, startTime, time)
	}
	if potential >= m.Threshold {
		spiked = true
		node.SetSpike(time, spiked)
	}
	if spiked {
		for _, syn := range node.GetSynapses() {
			queue.Add(time+core.Time(syn.GetDelay()), syn.GetTarget())
		}
	}
}

func (m *StdpKernel) Update(node core.Node, time core.Time) {
	// Apply STDP
}
