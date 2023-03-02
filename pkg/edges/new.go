package edges

import "spiky/pkg/core"

type baseEdge struct {
	source core.Node
	target core.Node
	weight float64
	delay  int
}

func (e *baseEdge) GetSource() core.Node {
	return e.source
}
func (e *baseEdge) GetTarget() core.Node {
	return e.target
}
func (e *baseEdge) GetWeight() float64 {
	return e.weight
}
func (e *baseEdge) GetDelay() int {
	return e.delay
}

func New(source core.Node, target core.Node) core.Edge {
	return &baseEdge{
		source: source,
		target: target,
		weight: 0.0,
		delay:  1,
	}
}
