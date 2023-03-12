package core

import (
	"spiky/pkg/utils"
)

type Model interface {
	GetInput() *Layer
	GetOutput() *Layer
	GetAllLayer() []*Layer
	GetLayer(idx int) *Layer
	Len() int
	Predict(input []byte) []byte
	Fit(input []byte, output []byte) []byte
	Clear()
}

type SampleModel struct {
	layers []*Layer
	codec  Codec
	world  *World
}

func NewSampleModel(codec Codec, layers []*Layer, constants *utils.Constants) *SampleModel {
	return &SampleModel{
		layers: layers,
		codec:  codec,
		world:  NewWorld(constants),
	}
}

func (model *SampleModel) Visit(fn func(neuron *Neuron)) {
	for i := 0; i < model.Len(); i++ {
		model.GetLayer(i).Visit(func(idx int, neuron *Neuron) {
			fn(neuron)
		})
	}
}

func (model *SampleModel) GetAllLayer() []*Layer {
	return model.layers
}

func (model *SampleModel) GetLayer(idx int) *Layer {
	return model.layers[idx]
}

func (model *SampleModel) GetInput() *Layer {
	return model.GetLayer(0)
}

func (model *SampleModel) GetOutput() *Layer {
	return model.layers[len(model.layers)-1]
}

func (model *SampleModel) Len() int {
	return len(model.layers)
}

func (model *SampleModel) Clear() {
	model.Visit(func(neuron *Neuron) {
		neuron.Clear()
	})
	model.world.Clear()
}

func (model *SampleModel) Predict(x []byte) []byte {
	input := model.GetInput()
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
	input.Visit(func(idx int, neuron *Neuron) {
		value := x[idx]
		spikes := model.codec.Encode(value)
		for _, time := range spikes {
			model.world.Schedule(time, neuron.Fire)
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
