package core

import (
	"spiky/pkg/utils"
	"testing"
)

func TestEdgeCreation(t *testing.T) {
	source := NewNeuron()
	target := NewNeuron()
	csts := utils.NewDefaultConstants()
	NewEdge(source, target, csts)
	if len(source.synapses) != 1 {
		t.Error("Invalid synapses count")
	}
	if len(target.dendrites) != 1 {
		t.Error("Invalid dendrites count")
	}
}
