package core

import (
	"errors"

	"github.com/aidarkhanov/nanoid/v2"
)

type Node interface {
	GetId() string
	Compute(time Time, queue *Queue)
	SetSpike(time Time, spiked bool)
	GetSpike(time Time) bool
	GetSpikeRate(startTime Time, endTime Time) (float64, error)
}

type BaseNode struct {
	id        string
	potential float64
	x         float64
	y         float64
	z         float64
	spikes    map[Time]bool

	synapses  []Edge
	dendrites []Edge
}

func (n *BaseNode) GetId() string {
	return n.id
}

func (n *BaseNode) Compute(time Time, queue *Queue) {
	// Compute new potential
	// Fire if superior to threshold
	if n.GetSpike(time) {
		for _, syn := range n.synapses {
			queue.Add(time+Time(syn.GetDelay()), syn.GetTarget())
		}
	}
}

func (n *BaseNode) GetSynapses() []Edge {
	return n.synapses
}

func (n *BaseNode) GetDendrites() []Edge {
	return n.dendrites
}

func (n *BaseNode) SetSpike(time Time, spiked bool) {
	n.spikes[time] = spiked
}

func (n *BaseNode) GetSpike(time Time) bool {
	return n.spikes[time]
}

func (n *BaseNode) GetSpikeRate(startTime Time, endTime Time) (float64, error) {
	if startTime >= endTime {
		return 0.0, errors.New("invalid time range")
	}
	var spikeCount float64 = 0.0
	for _, v := range n.spikes {
		if v {
			spikeCount++
		}
	}
	return spikeCount / float64(endTime-startTime), nil
}

func (n *BaseNode) GetChildren() []Node {
	var slice = make([]Node, len(n.synapses))
	for i, syn := range n.synapses {
		slice[i] = syn.GetTarget()
	}
	return slice
}

func (n *BaseNode) GetParents() []Node {
	var slice = make([]Node, len(n.dendrites))
	for i, syn := range n.dendrites {
		slice[i] = syn.GetSource()
	}
	return slice
}

func NewBaseNode() *BaseNode {
	id, _ := nanoid.New()
	node := BaseNode{
		id:        id,
		potential: 0,
		x:         0,
		y:         0,
		z:         0,
		spikes:    make(map[Time]bool),
		synapses:  []Edge{},
		dendrites: []Edge{},
	}
	return &node
}

func NewNodeSlice(size int) []Node {
	slice := make([]Node, size)
	for i := 0; i < size; i++ {
		node := NewBaseNode()
		slice[i] = node
	}
	return slice
}
