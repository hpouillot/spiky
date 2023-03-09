package core

import (
	"os"
	"reflect"
	"spiky/pkg/utils"
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
	csts := utils.NewDefaultConstants()
	output_size := 2
	codec := NewLatencyCodec(csts)
	input := NewLayer(2)
	output := NewLayer(output_size)
	DenseConnection(input, output, csts)
	model := NewSampleModel(codec, input, output, csts)
	prediction := model.Predict([]byte{255, 255})
	assert.Equal(t, len(prediction), output_size, "Invalid prediction size")
	assert.Equal(t, reflect.TypeOf(prediction[0]).Kind(), reflect.Uint8)
}
