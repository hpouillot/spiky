package data

import "spiky/pkg/core"

func NewNumberDataset(XSamples []float64, YSamples []float64) *Dataset {
	if len(XSamples) != len(YSamples) {
		panic("Samples must have the same size")
	}
	xSamples := make([][]float64, len(XSamples))
	ySamples := make([][]float64, len(YSamples))
	for idx, X := range XSamples {
		xSamples[idx] = []float64{X}
	}
	for idx, Y := range YSamples {
		ySamples[idx] = []float64{Y}
	}
	return &Dataset{
		get: func(idx int) core.Sample {
			return core.Sample{
				X: []float64{XSamples[idx]},
				Y: []float64{YSamples[idx]},
			}
		},
		len: len(xSamples),
		shape: Shape{
			X: 1,
			Y: 1,
		},
	}
}
