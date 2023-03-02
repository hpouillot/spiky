package nodes

import (
	"spiky/pkg/core"

	"github.com/aidarkhanov/nanoid/v2"
)

func Node(kernel core.Kernel) core.Node {
	id, _ := nanoid.New()
	node := baseNode{
		id:        id,
		potential: 0,
		position:  core.Point{},
		spikes:    make(map[core.Time]bool),
		synapses:  []core.Edge{},
		dendrites: []core.Edge{},
		kernel:    kernel,
	}
	return &node
}
