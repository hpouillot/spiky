package codec

import (
	"math"
	"spiky/pkg/utils"
)

func NewLatencyCodec(cst *utils.Constants) *LatencyCodec {
	return &LatencyCodec{
		constants: cst,
	}
}

type LatencyCodec struct {
	constants *utils.Constants
}

func (codec *LatencyCodec) Encode(value byte) []float64 {
	time := (float64(math.MaxUint8-value) / float64(math.MaxUint8)) * codec.constants.MaxTime
	spikes := []float64{time}
	return spikes
}

func (codec *LatencyCodec) Decode(spikes []float64) byte {
	firstSpikeTime := codec.constants.MaxTime
	for _, time := range spikes {
		firstSpikeTime = time
		break
	}
	return byte(((codec.constants.MaxTime - firstSpikeTime) / codec.constants.MaxTime) * math.MaxUint8)
}
