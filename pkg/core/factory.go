package core

import (
	"fmt"
	"spiky/pkg/codec"
)

func BuildSequential(layers []int, cfg *ModelConfig) *Model {
	codec := codec.NewLatencyCodec(255)
	modelLayers := []*Layer{}
	for idx, layerSize := range layers {
		newLayer := NewLayer(fmt.Sprintf("Layer_%d", idx), layerSize)
		modelLayers = append(modelLayers, newLayer)
		if idx >= 1 {
			DenseConnection(modelLayers[idx-1], modelLayers[idx], cfg)
		}
	}
	model := NewModel(codec, modelLayers, cfg)
	return model
}
