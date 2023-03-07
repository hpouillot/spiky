package core

import (
	"spiky/pkg/utils"
)

type Point interface {
	Coord() []uint64
	Shape() []uint64
}

type Connector interface {
	Connect(source Box[Neuron], target Box[Neuron])
}

type Model[I interface{}, O interface{}] interface {
	GetInput() Box[Neuron]
	GetOutput() Box[Neuron]
	Predict(input Box[I], duration float64) Box[O]
}

type SampleModel struct {
	input  Box[Neuron]
	output Box[Neuron]
	codec  Codec
	stack  *utils.TimeStack
}

func NewSampleModel(codec Codec, input Box[Neuron], output Box[Neuron]) SampleModel {
	return SampleModel{
		input:  input,
		output: output,
		codec:  codec,
		stack:  utils.NewTimeStack(),
	}
}

func (model *SampleModel) GetInput() Box[Neuron] {
	return model.input
}

func (model *SampleModel) GetOutput() Box[Neuron] {
	return model.output
}

func (model *SampleModel) Predict(x []byte, duration float64) []byte {
	input := model.GetInput()
	if input == nil {
		return []byte{}
	}
	input.Visit(func(idx int, node *Neuron) {
		value := x[idx]
		spikes := model.codec.Encode(value)
		for _, time := range spikes {
			model.stack.Push(time, node.Fire)
		}
	})
	model.stack.Resolve(0, duration)
	output := model.GetOutput()
	y := make([]byte, output.Size())
	output.Visit(func(idx int, node *Neuron) {
		y[idx] = model.codec.Decode(node.spikes)
	})
	return y
}
