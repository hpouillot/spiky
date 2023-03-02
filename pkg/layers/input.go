package layers

import (
	"spiky/pkg/core"
	"spiky/pkg/kernels"
)

func Input(source core.Dataset) core.Layer {
	layerSize := source.Size()
	kernel := &kernels.InputKernel{
		Dataset: source,
	}
	return Layer(layerSize, kernel)
}
