package kernels

import "spiky/pkg/core"

type InputKernel struct {
	Dataset core.Dataset
}

func (k *InputKernel) Compute(node core.Node, time core.Time, queue *core.Queue) {
	spiked := k.Dataset.Get(node.GetPosition(), time)
	if spiked {
		node.SetSpike(time, true)
		for _, syn := range node.GetSynapses() {
			queue.Add(time+core.Time(syn.GetDelay()), syn.GetTarget())
		}
	}
	queue.Add(time+1, node)
}

func (k *InputKernel) Update(node core.Node, time core.Time) {
	// Apply STDP
}
