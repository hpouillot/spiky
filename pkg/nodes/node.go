package nodes

import (
	"math/rand"
	"spiky/pkg/core"

	"github.com/aidarkhanov/nanoid/v2"
)

func Node(kernel core.Kernel) core.Node {
	id, _ := nanoid.New()
	node := baseNode{
		id:        id,
		potential: 0,
		position: core.Point{
			X: rand.Float64(),
			Y: rand.Float64(),
			Z: rand.Float64(),
		},
		spikes:    make(map[core.Time]bool),
		synapses:  []core.Edge{},
		dendrites: []core.Edge{},
		kernel:    kernel,
	}
	return &node
}
