package kernels

import "spiky/pkg/core"

type InputKernel struct {
	dataset core.Dataset
}

func (k *InputKernel) Compute(node core.Node, time core.Time) bool {
	return k.dataset.Get(node.GetPosition(), time)
}

func (k *InputKernel) Update(node core.Node, time core.Time) {
	// Apply STDP
}
