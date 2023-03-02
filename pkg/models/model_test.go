package models

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/edges"
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
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
	}) // Sized, Localized dataset ?
	target := data.Text([]string{
		"Y",
		"N",
		"Y",
		"N",
		"Y",
		"N",
	})

	input := layers.Input(source)
	output := layers.Output(target)

	edges.Dense(input, output, 0.5)

	model := Model(input, output)

	for k := 0; k < 5; k++ {
		model.Run(100)
		source.Next()
		target.Next()
	}
}

func TestModel(t *testing.T) {

}
