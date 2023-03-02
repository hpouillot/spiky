package layers

import (
	"spiky/pkg/core"
	"spiky/pkg/nodes"
)

type layer struct {
	nodes []core.Node
}

func (l *layer) GetNodes() []core.Node {
	return l.nodes
}

func (l *layer) Visit(fn func(node core.Node, idx int)) {
	for idx, node := range l.nodes {
		fn(node, idx)
	}
}

func (l *layer) GetNode(idx int) core.Node {
	return l.nodes[idx]
}

func (l *layer) Reset() {
	l.Visit(func(node core.Node, _ int) {
		node.Reset()
	})
}

func (l *layer) Size() int {
	return len(l.nodes)
}

func Layer(size int, kernel core.Kernel) core.Layer {
	layerNodes := make([]core.Node, size)
	for idx := range make([]int, size) {
		layerNodes[idx] = nodes.Node(kernel)
	}
	return &layer{
		nodes: layerNodes,
	}
}
