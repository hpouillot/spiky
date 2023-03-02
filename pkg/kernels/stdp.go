package kernels

import "spiky/pkg/core"

type StdpKernel struct{}

func (m *StdpKernel) Compute(node core.Node, time core.Time, queue *core.Queue) {
	if false {
		for _, syn := range node.GetSynapses() {
			queue.Add(time+core.Time(syn.GetDelay()), syn.GetTarget())
		}
	}
}

func (m *StdpKernel) Update(node core.Node, time core.Time) {
	// Apply STDP
}
