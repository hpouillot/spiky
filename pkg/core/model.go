package core

import (
	"math/rand"
	"spiky/pkg/utils"
)

type Model struct {
	layers []*Layer
	codec  ICodec
	World  *World
}

func NewModel(codec ICodec, layers []*Layer, world *World) *Model {
	return &Model{
		layers: layers,
		codec:  codec,
		World:  world,
	}
}

func (model *Model) VisitNeurons(visitor func(neuron *Neuron)) {
	for i := 0; i < model.Len(); i++ {
		model.GetLayer(i).Visit(func(idx int, neuron *Neuron) {
			visitor(neuron)
		})
	}
}

func (model *Model) VisitEdges(visitor func(edge IEdge)) {
	visitedEdges := map[string]IEdge{}
	var visitNeuron func(*Neuron)
	visitNeuron = func(node *Neuron) {
		for _, syn := range node.synapses {
			if _, ok := visitedEdges[syn.GetId()]; !ok {
				visitedEdges[syn.GetId()] = syn
				visitor(syn)
				visitNeuron(syn.GetTarget())
			}
		}
	}
	model.GetInput().Visit(func(idx int, neuron *Neuron) {
		visitNeuron(neuron)
	})
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
	model.codec.Encode(input, x)
}

func (model *Model) Decode() []float64 {
	output := model.GetOutput()
	return model.codec.Decode(output)
}

func (model *Model) DecodeClass() int {
	predictions := model.Decode()
	maxValue, _ := utils.Max(predictions)
	maxIndices := []int{}
	for idx, val := range predictions {
		if val == maxValue {
			maxIndices = append(maxIndices, idx)
		}
	}
	idx := rand.Intn(len(maxIndices))
	return maxIndices[idx]
}

func (model *Model) Run() {
	for model.World.Next() {
	}
}

func (model *Model) Stdp(reward float64) {
	model.VisitEdges(func(edge IEdge) {
		edge.Stdp(model.World, reward)
	})
}

func (model *Model) Fit(y []float64) {
	model.codec.Fit(model.GetOutput(), y)

}
