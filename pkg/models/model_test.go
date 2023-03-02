package models

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/edges"
	"spiky/pkg/kernels"
	"spiky/pkg/layers"
	"testing"
)

type MockedKernel struct {
	counter int
}

func (mk *MockedKernel) Compute(node *core.Node, time core.Time) bool {
	mk.counter += 1
	return true
}

func (mk *MockedKernel) Update(node *core.Node, time core.Time) {
	mk.counter += 1
}

func TestModelInstantiation(t *testing.T) {
	source := data.Text([]string{
		"🦾",
		"2",
		"3",
		"4",
		"5",
		"6",
	}) // Sized, Localized dataset ?

	kernel := kernels.StdpKernel{
		Threshold: 250.0,
		Tho:       10,
	}

	input := layers.Input(source)
	hidden := layers.Layer(100, &kernel)

	edges.Dense(input, hidden, 1.0)

	model := Model(input, hidden)

	for k := 0; k < 5; k++ {
		model.Run(100)
		source.Next(true)
	}
}

func TestModel(t *testing.T) {

}
