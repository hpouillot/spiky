package layers

import (
	"spiky/pkg/core"
	"spiky/pkg/kernels"
)

func Output(source core.Dataset) core.Layer {
	layerSize := source.Shape()
	kernel := &kernels.OutputKernel{}
	return Layer(layerSize[0], kernel)
}
