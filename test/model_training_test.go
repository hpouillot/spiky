package test

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/training"
	"spiky/pkg/utils"
	"testing"
)

func TestModelFitting(t *testing.T) {
	dataset := data.NewMnist("../../mnist")
	inputSize, outputSize := dataset.Shape()
	csts := utils.NewDefaultConstants()
	model := buildModel(inputSize, outputSize, csts)

	app := training.NewTrainingApp(model, dataset, csts)
	// defer app.Close()
	// app.Open()
	app.Start(10000)
}

func buildModel(inputSize int, outputSize int, csts *utils.Constants) core.Model {
	codec := core.NewLatencyCodec(csts)
	input := core.NewLayer("Input", inputSize)
	// hidden1 := NewLayer("Hidden 1", 100)
	// core.DenseConnection(input, hidden1, csts)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, output, csts)
	layers := []*core.Layer{
		input,
		// hidden1,
		output,
	}
	model := core.NewSampleModel(codec, layers, csts)
	return model
}
