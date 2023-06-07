package core

func DenseConnection(sourceLayer *Layer, targetLayer *Layer, config *ModelConfig) {
	sourceLayer.Visit(func(idx int, source *Neuron) {
		targetLayer.Visit(func(idx int, target *Neuron) {
			NewEdge(source, target, config)
		})
	})
}

func MutualConnection(layer *Layer, config *ModelConfig) {
	for i := 0; i < layer.Size(); i++ {
		for j := i + 1; j < layer.Size(); j++ {
			NewPositiveEdge(layer.Get(i), layer.Get(j), config)
			NewNegativeEdge(layer.Get(i), layer.Get(j), config)
		}
	}
}

func WTAConnection(layer *Layer, config *ModelConfig) {
	for i := 0; i < layer.Size(); i++ {
		for j := 0; j < layer.Size(); j++ {
			if i != j {
				NewNegativeEdge(layer.Get(i), layer.Get(j), config)
			}
		}
	}
}
