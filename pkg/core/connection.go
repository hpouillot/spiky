package core

import "spiky/pkg/utils"

func DenseConnection(sourceLayer *Layer, targetLayer *Layer, csts *utils.Constants) {
	sourceLayer.Visit(func(idx int, source *Neuron) {
		targetLayer.Visit(func(idx int, target *Neuron) {
			NewEdge(source, target, csts)
		})
	})
}
