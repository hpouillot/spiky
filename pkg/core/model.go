package core

import (
	"math"
	"spiky/pkg/utils"
)

type Model interface {
	GetInput() *Layer
	GetOutput() *Layer
	GetAllLayer() []*Layer
	GetLayer(idx int) *Layer
	Len() int
	Predict(input []byte) []byte
	Fit(input []byte, output []byte) ([]byte, float64)
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

func (model *SampleModel) Fit(x []byte, y []byte) ([]byte, float64) {
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
	predictions := make([]byte, output.Size())
	loss := 0.0
	output.Visit(func(idx int, node *Neuron) {
		spikes := model.codec.Encode(y[idx])
		lastSpike, err := node.GetLastSpikeTime()
		if err != nil {
			lastSpike = model.world.Const.MaxTime
		}
		node.Adjust(model.world, lastSpike-spikes[0])
		loss += math.Abs(lastSpike - spikes[0])
		predictions[idx] = model.codec.Decode(*node.GetSpikes())
	})
	return predictions, loss
}
