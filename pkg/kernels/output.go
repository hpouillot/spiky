package kernels

import "spiky/pkg/core"

type OutputKernel struct {
	Dataset core.Dataset
}

func (m *OutputKernel) GetMaxWeight() float64 {
	return 0
}

func (k *OutputKernel) Compute(node core.Node, time core.Time, queue core.Queue) {

}

func (k *OutputKernel) Update(node core.Node, time core.Time) {
	// Move sensors depending on dataset. Capture more data ?
}
