package test

import (
	"spiky/pkg/codec"
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/reporter"
	"spiky/pkg/utils"
	"testing"
)

func TestModelFitting(t *testing.T) {
	dataset := data.NewNumberDataset([]float64{100, 200}, []float64{1, 2})
	inputSize, outputSize := dataset.Shape()
	csts := utils.NewDefaultConstants()
	model := buildModel(inputSize, outputSize, csts)
	trainer := core.NewTrainer(model, dataset, csts)
	reporter.NewLogReporter(trainer)
	trainer.Start(1)
}

func buildModel(inputSize int, outputSize int, csts *utils.Constants) *core.Model {
	codec := codec.NewLatencyCodec(255, csts)
	input := core.NewLayer("Input", inputSize)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, output, csts)
	layers := []*core.Layer{
		input,
		output,
	}
	model := core.NewModel(codec, layers, csts)
	return model
}
