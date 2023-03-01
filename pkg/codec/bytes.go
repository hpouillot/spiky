package codec

import (
	"fmt"
	"spiky/pkg/core"

	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

type BytesCodec struct {
	length     int
	resolution int
}

func (codec BytesCodec) Size() int {
	return codec.length
}

func (codec BytesCodec) Encode(data []byte, nodes []core.Node) error {
	if len(data) > len(nodes) {
		return fmt.Errorf("sequence to encode %d is too long compared to number of nodes %d", len(data), len(nodes))
	}
	src := rand.New(rand.NewSource(1))
	for idx, currentByte := range data {
		node := nodes[idx]
		var spikeProba float64 = (float64(currentByte) / 255.0)
		distrib := distuv.Bernoulli{P: spikeProba, Src: src}
		for time := range make([]int, codec.resolution) {
			if distrib.Rand() == 1.0 {
				node.SetSpike(core.Time(time), true)
			} else {
				node.SetSpike(core.Time(time), false)
			}
		}
	}
	return nil
}

func (codec BytesCodec) Decode(nodes []core.Node) ([]byte, error) {
	bytesArray := make([]byte, len(nodes))
	for idx, node := range nodes {
		spikeRate, err := node.GetSpikeRate(0, 1)
		if err != nil {
			return []byte{}, nil
		}
		bytesArray[idx] = byte(spikeRate * 255)
	}
	return bytesArray, nil
}
