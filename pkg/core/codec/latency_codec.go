package codec

import (
	"math"
	"spiky/pkg/utils"
)

func NewLatencyCodec(maxValue float64, cst *utils.Constants) *LatencyCodec {
	return &LatencyCodec{
		maxValue:  maxValue,
		constants: cst,
	}
}

type LatencyCodec struct {
	maxValue  float64
	constants *utils.Constants
}

func (codec *LatencyCodec) Encode(value float64) []float64 {
	time := ((codec.maxValue - math.Min(value, codec.maxValue)) / codec.maxValue) * codec.constants.MaxTime
	spikes := []float64{time}
	return spikes
}

func (codec *LatencyCodec) Decode(spikes []float64) float64 {
	firstSpikeTime := codec.constants.MaxTime
	for _, time := range spikes {
		firstSpikeTime = time
		break
	}
	return ((codec.constants.MaxTime - firstSpikeTime) / codec.constants.MaxTime) * codec.maxValue
}
