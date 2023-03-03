package edges

import (
	"math/rand"
	"spiky/pkg/core"
)

type baseEdge struct {
	source core.Node
	target core.Node
	weight float64
	delay  float64
}

func (e *baseEdge) GetSource() core.Node {
	return e.source
}
func (e *baseEdge) GetTarget() core.Node {
	return e.target
}
func (e *baseEdge) GetDelay() float64 {
	return e.delay
}
func (e *baseEdge) GetWeight() float64 {
	return e.weight
}
func (e *baseEdge) UpdateWeight(deltaW float64) float64 {
	e.weight += deltaW
	return e.weight
}

func New(source core.Node, target core.Node) core.Edge {
	return &baseEdge{
		source: source,
		target: target,
		delay:  rand.Float64() * 2,
		weight: 0.0,
	}
}
