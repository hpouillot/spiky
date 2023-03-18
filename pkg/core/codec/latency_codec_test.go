package codec

import (
	"fmt"
	"spiky/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatencyCodec(t *testing.T) {
	constants := utils.NewDefaultConstants()
	codec := NewLatencyCodec(255, constants)
	spikes := codec.Encode(155)
	if len(spikes) > 1 {
		t.Error("Invalid spike count")
	}
	fmt.Println(spikes)
	assert.GreaterOrEqual(t, spikes[0], 0.0)
	assert.LessOrEqual(t, spikes[0], 10.0)

	spikesToDecode := []float64{2.1555}
	value := codec.Decode(spikesToDecode)
	// fmt.Print(value)
	assert.GreaterOrEqual(t, value, float64(100))
	assert.LessOrEqual(t, value, float64(200))
}
