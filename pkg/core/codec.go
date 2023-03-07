package core

import (
	"math"
	"sort"
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

// Time to first spike
type RateCodec struct {
	duration float64
}

func NewRateCodec(duration float64) *RateCodec {
	return &RateCodec{
		duration: duration,
	}
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
	// Compute number of events
	// Take that many rand numbers in interval
	return spikes
}

func (codec *RateCodec) Decode(spikes []float64) byte {
	sppikeCount := math.Max(float64(len(spikes)), codec.duration)
	rate := sppikeCount / codec.duration
	// Get poisson
	return uint8(rate * math.MaxUint8)
}
