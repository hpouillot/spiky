package core

func DenseConnection(sourceLayer *Layer, targetLayer *Layer, config *ModelConfig) {
	sourceLayer.Visit(func(idx int, source *Neuron) {
		targetLayer.Visit(func(idx int, target *Neuron) {
			NewEdge(source, target, config)
		})
	})
}

func MutualConnection(layer *Layer, config *ModelConfig) {
	layer.Visit(func(idx int, source *Neuron) {
		layer.Visit(func(idx int, target *Neuron) {
			if source != target {
				NewEdge(source, target, config)
			}
		})
	})
}
