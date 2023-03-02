package env

import (
	"math"
	"spiky/pkg/core"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type byteDataset struct {
	distributions []distuv.Bernoulli
	samples       [][]byte
	size          int
	cursor        int
}

func (bd *byteDataset) Get(position core.Point, time core.Time) bool {
	idx := int(position.X * float64(bd.size))
	if idx > bd.size {
		panic("invalid access")
	}
	sample := bd.distributions[idx].Rand()
	if sample == 1 {
		return true
	} else {
		return false
	}
}

func (bd *byteDataset) Next() {
	bd.cursor += 1
	values := bd.samples[bd.cursor]
	for idx, value := range values {
		src := rand.New(rand.NewSource(1))
		spikeProba := float64(value) / math.MaxUint8
		bd.distributions[idx] = distuv.Bernoulli{P: spikeProba, Src: src}
	}
}

func (bd *byteDataset) HasNext() bool {
	if bd.cursor == len(bd.samples)-1 {
		return false
	} else {
		return true
	}
}

func (bd *byteDataset) Reset() {
	bd.cursor = 0
	bd.distributions = make([]distuv.Bernoulli, bd.size)
}

func (bd *byteDataset) Size() int {
	return bd.size
}
