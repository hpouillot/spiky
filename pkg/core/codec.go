package core

import (
	"math"
	"sort"
	"spiky/pkg/utils"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type Encoder interface {
	Encode(value byte) []float64 // Schedule spikes for nodes
}

type Decoder interface {
	Decode(spikes []float64) byte
}

type Codec interface {
	Encoder
	Decoder
}

func NewRateCodec(csts *utils.Constants) *RateCodec {
	return &RateCodec{
		duration: csts.MaxTime,
	}
}

func NewLatencyCodec(cst *utils.Constants) *LatencyCodec {
	return &LatencyCodec{
		constants: cst,
	}
}

// Time to first spike
type RateCodec struct {
	duration float64
}

func (codec *RateCodec) Encode(value byte) []float64 {
	lambda := float64(value) / codec.duration
	src := rand.NewSource(uint64(time.Now().Unix()))
	distrib := distuv.Poisson{
		Lambda: lambda,
		Src:    src,
	}
	eventsCount := int(distrib.Rand())
	spikes := make([]float64, eventsCount)
	r := rand.New(src)
	for i := 0; i < eventsCount; i++ {
		spikes[i] = r.Float64() * codec.duration
	}
	sort.Float64s(spikes)
	return spikes
}

func (codec *RateCodec) Decode(spikes []float64) byte {
	sppikeCount := math.Max(float64(len(spikes)), codec.duration)
	rate := sppikeCount / codec.duration
	// Get poisson
	return uint8(rate * math.MaxUint8)
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
