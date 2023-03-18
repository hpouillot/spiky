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

func (codec *LatencyCodec) Encode(value *float64) *float64 {
	if value == nil || *value == 0 {
		return nil
	}
	time := -(math.Log(*value/codec.maxValue) * codec.tho)
	return &time
}

func (codec *LatencyCodec) Decode(time *float64) *float64 {
	if time == nil {
		return new(float64)
	}
	value := math.Exp(-*time/codec.tho) * codec.maxValue
	return &value
}
