package test

import (
	"spiky/pkg/codec"
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/reporter"
	"testing"
)

func TestModelFitting(t *testing.T) {
	dataset := data.NewNumberDataset([]float64{100, 200}, []float64{1, 2})
	inputSize, outputSize := dataset.Shape()
	config := core.NewDefaultConfig()
	model := buildModel(inputSize, outputSize, config)
	trainer := core.NewTrainer(model, dataset)
	reporter.NewLogReporter(trainer)
	trainer.Start(1)
}

func buildModel(inputSize int, outputSize int, config *core.ModelConfig) *core.Model {
	codec := codec.NewLatencyCodec(255)
	input := core.NewLayer("Input", inputSize)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, output, config)
	layers := []*core.Layer{
		input,
		output,
	}
	model := core.NewModel(codec, layers, config)
	return model
}
