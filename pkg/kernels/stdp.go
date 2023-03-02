package kernels

import "spiky/pkg/core"

type StdpKernel struct{}

func (m *StdpKernel) Compute(node core.Node, time core.Time) bool {
	return false
}

func (m *StdpKernel) Update(node core.Node, time core.Time) {
	// Apply STDP
}
