package core

import "testing"

func TestEdgeCreation(t *testing.T) {
	source := NewNeuron(10, 0)
	target := NewNeuron(10, 0)
	NewEdge(source, target)
	if len(source.synapses) != 1 {
		t.Error("Invalid synapses count")
	}
	if len(target.dendrites) != 1 {
		t.Error("Invalid dendrites count")
	}
}
