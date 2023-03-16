package core

import (
	"math"
	"spiky/pkg/utils"
)

type IModel interface {
	GetInput() *Layer
	GetOutput() *Layer
	GetAllLayer() []*Layer
	GetLayer(idx int) *Layer
	Len() int
	Encode(input []float64)
	Decode() []float64
	Run()
	Adjust(output []float64) float64
	Reset()
}

type SampleModel struct {
	layers []*Layer
	codec  ICodec
	world  *World
}

func NewSampleModel(codec ICodec, layers []*Layer, constants *utils.Constants) *SampleModel {
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

func (model *SampleModel) Reset() {
	model.world.Reset()
}

func (model *SampleModel) Encode(x []float64) {
	input := model.GetInput()
	input.Visit(func(idx int, node *Neuron) {
		spikes := model.codec.Encode(x[idx])
		for _, time := range spikes {
			model.world.Schedule(time, node.Fire)
		}
	})
}

func (model *SampleModel) Decode() []float64 {
	output := model.GetOutput()
	y := make([]float64, output.Size())
	output.Visit(func(idx int, node *Neuron) {
		y[idx] = model.codec.Decode(node.spikes)
	})
	return y
}

func (model *SampleModel) Run() {
	for model.world.Next() {
	}
}

func (model *SampleModel) Adjust(y []float64) float64 {
	loss := 0.0
	model.GetOutput().Visit(func(idx int, node *Neuron) {
		expectedSpikes := model.codec.Encode(y[idx])
		lastSpike, err := node.GetLastSpikeTime()
		if err != nil {
			lastSpike = model.world.Const.MaxTime
		}
		lastExpectedSpike := model.world.Const.MaxTime
		if len(expectedSpikes) != 0 {
			lastExpectedSpike = expectedSpikes[0]
		}
		node.Adjust(model.world, lastSpike-lastExpectedSpike)
		loss += math.Abs(lastSpike - lastExpectedSpike)
	})
	return loss
}
