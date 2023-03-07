package core

import (
	"fmt"
	"testing"
)

func TestModelCreation(t *testing.T) {
	codec := NewRateCodec(10)
	input := NewLayer(2)
	output := NewLayer(2)
	input.Visit(func(idx int, source *Neuron) {
		output.Visit(func(idx int, target *Neuron) {
			NewEdge(source, target)
		})
	})
	model := NewSampleModel(codec, input, output)
	outputValue := model.Predict([]byte{255, 255}, 10)
	fmt.Println(outputValue)
}
