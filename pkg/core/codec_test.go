package core

import (
	"fmt"
	"spiky/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateCodec(t *testing.T) {
	constants := utils.NewDefaultConstants()
	codec := NewRateCodec(constants)
	spikes := codec.Encode(255)
	if len(spikes) < 5 {
		t.Error("Invalid spike count")
	}
}

func TestLatencyCodec(t *testing.T) {
	constants := utils.NewDefaultConstants()
	codec := NewLatencyCodec(constants)
	spikes := codec.Encode(255)
	if len(spikes) > 1 {
		t.Error("Invalid spike count")
	}

	spikesToDecode := []float64{15.1555}
	value := codec.Decode(spikesToDecode)
	fmt.Print(value)
	assert.GreaterOrEqual(t, value, byte(100))
	assert.LessOrEqual(t, value, byte(200))
}
