package layers

import (
	"spiky/pkg/core"
	"spiky/pkg/kernels"
)

func Output(source core.Dataset) core.Layer {
	layerSize := source.Size()
	kernel := &kernels.StdpKernel{}
	return Layer(layerSize, kernel)
}
