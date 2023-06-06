package core

import (
	"fmt"
	"math/rand"
	"spiky/pkg/utils"
)

type Model struct {
	layers []*Layer
	codec  ICodec
	World  *World
}

func NewModel(codec ICodec, layers []*Layer, cfg *ModelConfig) *Model {
	return &Model{
		layers: layers,
		codec:  codec,
		World:  NewWorld(cfg),
	}
}

func (model *Model) VisitNeurons(visitor func(neuron *Neuron)) {
	for i := 0; i < model.Len(); i++ {
		model.GetLayer(i).Visit(func(idx int, neuron *Neuron) {
			visitor(neuron)
		})
	}
}

func (model *Model) VisitEdges(visitor func(edge *Edge)) {
	visitedEdges := map[string]*Edge{}
	var visitNeuron func(*Neuron)
	visitNeuron = func(node *Neuron) {
		for _, syn := range node.synapses {
			if _, ok := visitedEdges[syn.source.id+syn.target.id]; !ok {
				visitedEdges[syn.source.id+syn.target.id] = syn
				visitor(syn)
				visitNeuron(syn.target)
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

func (model *Model) DecodeClass() int {
	predictions := model.Decode()
	fmt.Printf("Predictions %v", predictions)
	maxValue := utils.Max(predictions)
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
	model.VisitEdges(func(edge *Edge) {
		edge.Stdp(model.World, reward)
	})
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
