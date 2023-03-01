package codec

import (
	"errors"
	"spiky/pkg/core"
)

type BinaryCodec struct {
	resolution int
}

func (bc BinaryCodec) Size() int {
	return 1
}

func (bc BinaryCodec) Encode(data []bool, nodes []core.Node) error {
	if len(nodes) != bc.Size() {
		errors.New("Invalid nodes length")
	}
	for _, value := range data {
		for time := range make([]int, bc.resolution) {
			nodes[0].SetSpike(core.Time(time), value)
		}
	}
	return nil
}

func (bc BinaryCodec) Decode(nodes []core.Node) (bool, error) {
	// spikeRate := nodes[0].GetSpikeRate()
	// if spikeRate > 0.5 {
	// 	return true, nil
	// } else {
	// 	return false, nil
	// }
	return false, nil
}
