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
		spikes := model.codec.Encode(x[idx])
		for _, time := range spikes {
			model.World.Schedule(time, node.Fire)
		}
	})
}

func (model *Model) Decode() []float64 {
	output := model.GetOutput()
	y := make([]float64, output.Size())
	output.Visit(func(idx int, node *Neuron) {
		y[idx] = model.codec.Decode(node.spikes)
	})
	return y
}

func (model *Model) Run() {
	for model.World.Next() {
	}
}

func (model *Model) Adjust(y []float64) float64 {
	model.GetOutput().Visit(func(idx int, node *Neuron) {
		expectedSpikes := model.codec.Encode(y[idx])
		lastSpike, err := node.GetLastSpikeTime()
		if err != nil {
			lastSpike = model.World.Const.MaxTime
		}
		lastExpectedSpike := model.World.Const.MaxTime
		if len(expectedSpikes) != 0 {
			lastExpectedSpike = expectedSpikes[0]
		}

		node.Adjust(model.World, lastSpike-lastExpectedSpike)
	})
	return 0.0
}
