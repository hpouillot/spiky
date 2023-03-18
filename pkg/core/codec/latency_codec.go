package codec

import (
	"math"
	"spiky/pkg/utils"
)

func NewLatencyCodec(maxValue float64, cst *utils.Constants) *LatencyCodec {
	return &LatencyCodec{
		tho:       1.5,
		maxValue:  maxValue,
		constants: cst,
	}
}

type LatencyCodec struct {
	tho       float64
	maxValue  float64
	constants *utils.Constants
}

func (codec *LatencyCodec) Encode(value float64) []float64 {
	if value == 0 {
		return []float64{}
	}
	time := -(math.Log(value/codec.maxValue) * codec.tho)
	// ((codec.maxValue - math.Min(value, codec.maxValue)) / codec.maxValue) * codec.constants.MaxTime
	spikes := []float64{time}
	return spikes
}

func (codec *LatencyCodec) Decode(spikes []float64) float64 {
	firstSpikeTime := codec.constants.MaxTime
	for _, time := range spikes {
		firstSpikeTime = time
		break
	}
	value := math.Exp(-firstSpikeTime/codec.tho) * codec.maxValue
	return value
	// return ((codec.constants.MaxTime - math.Min(firstSpikeTime, codec.constants.MaxTime)) / codec.constants.MaxTime) * codec.maxValue
}
