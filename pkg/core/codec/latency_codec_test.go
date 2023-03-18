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
	value := 155.0
	time := codec.Encode(&value)
	fmt.Println(time)
	assert.GreaterOrEqual(t, *time, 0.0)
	assert.LessOrEqual(t, *time, 10.0)

	timeToDecode := 2.1555
	decodedValue := codec.Decode(&timeToDecode)
	assert.GreaterOrEqual(t, *decodedValue, 50.0)
	assert.LessOrEqual(t, *decodedValue, 200.0)
}
