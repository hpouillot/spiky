package factory

import (
	"fmt"
	"spiky/pkg/codec"
	"spiky/pkg/core"
)

func BuildSequential(world *core.World, layers []int) *core.Model {
	codec := codec.NewLatencyCodec(world, 255)
	modelLayers := []*core.Layer{}
	for idx, layerSize := range layers {
		newLayer := core.NewLayer(fmt.Sprintf("Layer_%d", idx), layerSize)
		modelLayers = append(modelLayers, newLayer)
		if idx >= 1 {
			core.DenseConnection(modelLayers[idx-1], modelLayers[idx], world.Const)
		}
	}
	model := core.NewModel(codec, modelLayers, world)
	return model
}
