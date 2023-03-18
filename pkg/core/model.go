package core

import (
	"spiky/pkg/utils"
)

type Model struct {
	layers []*Layer
	codec  ICodec
	World  *World
}

func NewModel(codec ICodec, layers []*Layer, constants *utils.Constants) *Model {
	return &Model{
		layers: layers,
		codec:  codec,
		World:  NewWorld(constants),
	}
}

func (model *Model) Visit(fn func(neuron *Neuron)) {
	for i := 0; i < model.Len(); i++ {
		model.GetLayer(i).Visit(func(idx int, neuron *Neuron) {
			fn(neuron)
		})
	}
}

func (model *Model) GetAllLayer() []*Layer {
	return model.layers
}

func (model *Model) GetLayer(idx int) *Layer {
	return model.layers[idx]
}

func (model *Model) GetInput() *Layer {
	return model.GetLayer(0)
}

func (model *Model) GetOutput() *Layer {
	return model.layers[len(model.layers)-1]
}

func (model *Model) Len() int {
	return len(model.layers)
}

func (model *Model) Reset() {
	model.World.Reset()
}

func (model *Model) Encode(x []float64) {
	input := model.GetInput()
	input.Visit(func(idx int, node *Neuron) {
		spikeTime := model.codec.Encode(&x[idx])
		if spikeTime != nil {
			model.World.Schedule(*spikeTime, node.Fire)
		}
	})
}

func (model *Model) Decode() []float64 {
	output := model.GetOutput()
	y := make([]float64, output.Size())
	output.Visit(func(idx int, node *Neuron) {
		y[idx] = *model.codec.Decode(node.GetSpikeTime())
	})
	return y
}

func (model *Model) Run() {
	for model.World.Next() {
	}
}

func (model *Model) Adjust(y []float64) float64 {
	model.GetOutput().Visit(func(idx int, node *Neuron) {
		expectedSpikeTime := model.codec.Encode(&y[idx])
		spikeTime := node.GetSpikeTime()
		if spikeTime == nil {
			spikeTime = &model.World.Const.MaxTime
		}
		if expectedSpikeTime == nil {
			node.Adjust(model.World, *spikeTime-model.World.Const.MaxTime)
		} else {
			node.Adjust(model.World, *spikeTime-*expectedSpikeTime)
		}
	})
	return 0.0
}
