package test

import (
	"os"
	"reflect"
	"spiky/pkg/codec"
	"spiky/pkg/core"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.ErrorLevel)
	code := m.Run()
	os.Exit(code)
}

func TestModelCreation(t *testing.T) {
	csts := core.NewDefaultConfig()
	output_size := 2
	world := core.NewWorld(csts)
	codec := codec.NewLatencyCodec(world, 255)
	input := core.NewLayer("Input", 2)
	output := core.NewLayer("Output", output_size)
	core.DenseConnection(input, output, csts)
	model := core.NewModel(codec, []*core.Layer{input, output}, world)
	model.Encode([]float64{255, 255})
	model.Run()
	prediction := model.Decode()
	assert.Equal(t, len(prediction), output_size, "Invalid prediction size")
	assert.Equal(t, reflect.TypeOf(prediction[0]).Kind(), reflect.Float64)
}

func TestModelVisit(t *testing.T) {
	csts := core.NewDefaultConfig()
	input := core.NewLayer("Input", 2)
	hidden := core.NewLayer("Hidden", 10)
	output := core.NewLayer("Output", 10)
	world := core.NewWorld(csts)
	codec := codec.NewLatencyCodec(world, 255)
	core.DenseConnection(input, hidden, csts)
	core.DenseConnection(hidden, output, csts)
	model := core.NewModel(codec, []*core.Layer{input, hidden, output}, world)
	neuronsCount := 0
	edgesCount := 0
	model.VisitNeurons(func(n *core.Neuron) {
		neuronsCount++
	})
	model.VisitEdges(func(e core.IEdge) {
		edgesCount++
	})
	assert.Equal(t, 22, neuronsCount, "Invalid neuron visit")
	assert.Equal(t, 120, edgesCount, "Invalid neuron visit")
}
