package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeCreation(t *testing.T) {
	source := NewNeuron("1")
	target := NewNeuron("2")
	csts := NewDefaultConfig()
	NewPositiveEdge(source, target, csts)
	assert.Equal(t, len(source.synapses), 1, "Invalid synapses count")
	assert.Equal(t, len(target.dendrites), 1, "Invalid dendrites count")
}
