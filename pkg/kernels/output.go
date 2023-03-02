package kernels

import "spiky/pkg/core"

type OutputKernel struct {
	Dataset core.Dataset
}

func (k *OutputKernel) Compute(node core.Node, time core.Time, queue *core.Queue) {
	spiked := k.Dataset.Get(node.GetPosition(), time)
	if spiked {
		node.SetSpike(time, true)
	}
}

func (k *OutputKernel) Update(node core.Node, time core.Time) {
	// Move sensors depending on dataset. Capture more data ?
}
