package data

import (
	"spiky/pkg/core"
	"spiky/pkg/utils"

	"gonum.org/v1/gonum/stat/distuv"
)

func Text(samples []string) core.Dataset {
	bytesSamples := make([][]byte, len(samples))
	maxLength := 0
	for idx, text := range samples {
		bytesSamples[idx] = []byte(text)
		maxLength = utils.MaxInt(maxLength, len(bytesSamples[idx]))
	}
	dataset := byteDataset{
		samples:       bytesSamples,
		cursor:        -1,
		size:          maxLength,
		distributions: make([]distuv.Bernoulli, maxLength),
	}
	dataset.Next(true)
	return &dataset
}
