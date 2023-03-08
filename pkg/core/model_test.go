package core

import (
	"spiky/pkg/utils"
	"testing"
)

func TestModelCreation(t *testing.T) {
	codec := NewRateCodec(10)
	input := NewLayer(2)
	output := NewLayer(2)
	constants := &utils.Constants{
		MaxWeight:        20,
		MinWeight:        0,
		LearningRate:     0.1,
		Threshold:        200.0,
		RefractoryPeriod: 1.0,
		Tho:              5.0,
		MaxDelay:         1.0,
	}
	input.Visit(func(idx int, source *Neuron) {
		output.Visit(func(idx int, target *Neuron) {
			NewEdge(source, target, constants)
		})
	})
	model := NewSampleModel(codec, input, output, constants)
	outputValue := model.Predict([]byte{255, 255}, 10)
	if len(outputValue) != 2 {
		t.Error("Invalid output size")
	}
}
