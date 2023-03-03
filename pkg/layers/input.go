package layers

import (
	"spiky/pkg/core"
	"spiky/pkg/kernels"
)

func Input(source core.Dataset, period core.Time) core.Layer {
	layerSize := source.Shape()
	kernel := &kernels.InputKernel{
		Dataset: source,
		Period:  period,
	}
	return Layer(layerSize[0], kernel)
}
