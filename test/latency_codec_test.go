package test

import (
	"fmt"
	"spiky/pkg/codec"
	"spiky/pkg/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatencyCodec(t *testing.T) {
	world := core.NewWorld(&core.ModelConfig{})
	codec := codec.NewLatencyCodec(world, 255)
	value := 155.0
	time := codec.ValueToTime(&value)
	fmt.Println(time)
	assert.GreaterOrEqual(t, *time, 0.0)
	assert.LessOrEqual(t, *time, 10.0)

	timeToDecode := 2.1555
	decodedValue := codec.TimeToValue(&timeToDecode)
	assert.GreaterOrEqual(t, decodedValue, 50.0)
	assert.LessOrEqual(t, decodedValue, 200.0)
}
