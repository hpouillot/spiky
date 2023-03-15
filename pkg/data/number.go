package data

import "spiky/pkg/core"

func NewNumberDataset(XSamples []byte, YSamples []byte) *Dataset {
	if len(XSamples) != len(YSamples) {
		panic("Samples must have the same size")
	}
	xSamples := make([][]byte, len(XSamples))
	ySamples := make([][]byte, len(YSamples))
	for idx, X := range XSamples {
		xSamples[idx] = []byte{X}
	}
	for idx, Y := range YSamples {
		ySamples[idx] = []byte{Y}
	}
	return &Dataset{
		get: func(idx int) core.Sample {
			return core.Sample{
				X: []byte{XSamples[idx]},
				Y: []byte{YSamples[idx]},
			}
		},
		len: len(xSamples),
		shape: Shape{
			X: 1,
			Y: 1,
		},
	}
}
