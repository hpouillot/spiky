package test

import (
	"spiky/pkg/core"
	"spiky/pkg/core/codec"
	"spiky/pkg/data"
	"spiky/pkg/utils"
	"testing"
)

func TestModelFitting(t *testing.T) {
	dataset := data.NewMnist("/Users/huguespouillot/go/src/spiky/mnist")
	inputSize, outputSize := dataset.Shape()
	csts := utils.NewDefaultConstants()
	model := buildModel(inputSize, outputSize, csts)
	trainer := core.NewTrainer(model, dataset, csts)
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
