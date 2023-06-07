package codec

import (
	"math"
	"spiky/pkg/core"
)

func NewLatencyCodec(world *core.World, maxValue float64) *LatencyCodec {
	return &LatencyCodec{
		world:    world,
		tho:      1.5,
		maxValue: maxValue,
	}
}

type LatencyCodec struct {
	world    *core.World
	tho      float64
	maxValue float64
}

func (codec *LatencyCodec) Encode(layer *core.Layer, input []float64) {
	layer.Visit(func(idx int, node *core.Neuron) {
		spikeTime := codec.ValueToTime(&input[idx])
		node.SetSpikeTime(codec.world, spikeTime)
		if spikeTime != nil {
			codec.world.Schedule(*spikeTime, node.Fire)
		}
	})
}

func (codec *LatencyCodec) Decode(layer *core.Layer) []float64 {
	y := make([]float64, layer.Size())
	layer.Visit(func(idx int, node *core.Neuron) {
		y[idx] = codec.TimeToValue(node.GetSpikeTime())
	})
	return y
}

func (codec *LatencyCodec) Fit(layer *core.Layer, output []float64) {
	layer.Visit(func(idx int, node *core.Neuron) {
		expectedSpikeTime := codec.ValueToTime(&output[idx])
		spikeTime := node.GetSpikeTime()
		if spikeTime == nil {
			spikeTime = &codec.world.Const.MaxTime
		}
		if expectedSpikeTime == nil {
			expectedSpikeTime = &codec.world.Const.MaxTime
		}
		node.Adjust(codec.world, *spikeTime-*expectedSpikeTime)
	})
}

func (codec *LatencyCodec) ValueToTime(value *float64) *float64 {
	time := new(float64)
	if value == nil || *value <= 0 {
		return nil
	} else if *value >= codec.maxValue {
		*time = 0.0
		return time
	} else {
		*time = -(math.Log(*value/codec.maxValue) * codec.tho)
		return time
	}
}

func (codec *LatencyCodec) TimeToValue(time *float64) float64 {
	if time == nil {
		return 0
	}
	value := math.Exp(-*time/codec.tho) * codec.maxValue
	return value
}
