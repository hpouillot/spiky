package edges

import (
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

func (e *baseEdge) SetWeight(weight float64) float64 {
	e.weight = weight
	return e.weight
}

func New(source core.Node, target core.Node, weight float64, delay float64) core.Edge {
	edge := &baseEdge{
		source: source,
		target: target,
		delay:  delay,
		weight: weight,
	}
	source.AddSynapse(edge)
	target.AddDendrite(edge)
	return edge
}
