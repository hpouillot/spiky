package kernels

import (
	"math"
	"spiky/pkg/core"
)

type StdpKernel struct {
	Threshold      float64
	Tho            float64
	LearningRate   float64
	MaxWeight      float64
	RefractoryTime core.Time
	TraceTarget    float64
	MaxDelay       float64
}

func (m *StdpKernel) GetMaxWeight() float64 {
	return m.MaxWeight
}

func (m *StdpKernel) GetContrib(parent core.Node, startTime core.Time, endTime core.Time) float64 {
	contrib := 0.0

	for _, time := range parent.GetSpikeTimes(startTime, endTime) {
		contrib += math.Exp(-(float64(endTime - time)) / m.Tho)
	}

	return contrib
}

func (m *StdpKernel) Compute(node core.Node, time core.Time, queue core.Queue) {
	potential := 0.0
	spiked := false
	startTime := node.GetLastSpikeTime()
	if time < startTime+m.RefractoryTime {
		return
	}
	for _, dendrite := range node.GetDendrites() {
		contrib := m.GetContrib(dendrite.GetSource(), startTime, time)
		potential += dendrite.GetWeight() * contrib
	}

	if potential >= m.Threshold {
		node.SetSpike(time, spiked)
		for _, syn := range node.GetSynapses() {
			queue.Add(time+core.Time(syn.GetDelay()), syn.GetTarget())
		}
		for _, dendrite := range node.GetDendrites() {
			weight := dendrite.GetWeight()
			contrib := m.GetContrib(dendrite.GetSource(), startTime, time)
			deltaW := m.LearningRate * (contrib - m.TraceTarget)
			if deltaW > 0 {
				dendrite.SetWeight(weight + deltaW*(m.MaxWeight-dendrite.GetWeight()))
			} else if deltaW < 0 {
				dendrite.SetWeight(weight + deltaW*dendrite.GetWeight())
			}
		}
	}
}

func (m *StdpKernel) Update(node core.Node, time core.Time) {
	// Apply STDP
}
