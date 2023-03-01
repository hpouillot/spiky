package codec

import (
	"errors"
	"spiky/pkg/core"
)

type CategoryCodec struct {
	binaryCodec  BinaryCodec
	size         int
	includeOther bool
}

func (cc CategoryCodec) GetNodeSliceFotCategory(category int, nodes []core.Node) []core.Node {
	return nodes[category : category+1]
}

func (cc CategoryCodec) Size() int {
	if cc.includeOther {
		return cc.size + 1
	} else {
		return cc.size
	}
}

func (cc CategoryCodec) Encode(data []int, nodes []core.Node) error {
	nodesLen := len(nodes)
	if nodesLen != cc.Size() {
		return errors.New("Invalid nodes size")
	}
	// for idx, value := range data {
	// if value > cc.size {
	// 	// Other
	// 	cc.binaryCodec.Encode(true, cc.GetNodeSliceFotCategory(len(nodes)-1, nodes))
	// } else {
	// 	cc.binaryCodec.Encode(true, cc.GetNodeSliceFotCategory(data, nodes))
	// }
	// }
	return nil
}

func (cc CategoryCodec) Decode(nodes []core.Node) (int, error) {
	maxRate := 0.0
	maxClass := 0
	for idx, node := range nodes {
		nodeRate, err := node.GetSpikeRate(0, 10)
		if err != nil {
			return 0, err
		}
		if nodeRate > maxRate {
			maxRate = nodeRate
			maxClass = idx
		}
	}
	return maxClass, nil
}
