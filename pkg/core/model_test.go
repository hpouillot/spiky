package core

import "testing"

type MockedKernel struct {
	counter int
}

func (mk *MockedKernel) Compute(node *Node, time Time) bool {
	mk.counter += 1
	return true
}

func (mk *MockedKernel) Update(node *Node, time Time) {
	mk.counter += 1
}

func TestModelInstantiation(t *testing.T) {
	Inputs := NewNodeSlice(10)
	Outputs := NewNodeSlice(20)

	model := Model{
		Inputs: Inputs,
		Outputs: Outputs,
	}

	model.Run(10)
}

func TestModel(t *testing.T) {

}
