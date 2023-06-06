package codec

import (
	"math"
)

func NewLatencyCodec(maxValue float64) *LatencyCodec {
	return &LatencyCodec{
		tho:      1.5,
		maxValue: maxValue,
	}
}

type LatencyCodec struct {
	tho      float64
	maxValue float64
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
