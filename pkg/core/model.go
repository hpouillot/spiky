package core

import (
	"spiky/pkg/utils"
)

type Model interface {
	GetInput() Box[Neuron]
	GetOutput() Box[Neuron]
	Predict(input []byte) []byte
	Fit(input []byte, output []byte) []byte
}

type SampleModel struct {
	input  Box[Neuron]
	output Box[Neuron]
	codec  Codec
	world  *World
}

func NewSampleModel(codec Codec, input Box[Neuron], output Box[Neuron], constants *utils.Constants) *SampleModel {
	return &SampleModel{
		input:  input,
		output: output,
		codec:  codec,
		world:  NewWorld(constants),
	}
}

func (model *SampleModel) GetInput() Box[Neuron] {
	return model.input
}

func (model *SampleModel) GetOutput() Box[Neuron] {
	return model.output
}

func (model *SampleModel) Predict(x []byte) []byte {
	input := model.GetInput()
	if input == nil {
		return []byte{}
	}
	input.Visit(func(idx int, node *Neuron) {
		value := x[idx]
		spikes := model.codec.Encode(value)
		for _, time := range spikes {
			model.world.Schedule(time, node.Fire)
		}
	})
	for model.world.Next() {
	}
	output := model.GetOutput()
	y := make([]byte, output.Size())
	output.Visit(func(idx int, node *Neuron) {
		y[idx] = model.codec.Decode(node.spikes)
	})
	return y
}

func (model *SampleModel) Fit(x []byte, y []byte) []byte {
	input := model.GetInput()
	if input == nil {
		return []byte{}
	}
	input.Visit(func(idx int, node *Neuron) {
		value := x[idx]
		spikes := model.codec.Encode(value)
		for _, time := range spikes {
			model.world.Schedule(time, node.Fire)
		}
	})
	for model.world.Next() {
	}
	output := model.GetOutput()
	prediction := make([]byte, output.Size())
	output.Visit(func(idx int, node *Neuron) {
		prediction[idx] = model.codec.Decode(node.spikes)
	})
	return prediction
}
