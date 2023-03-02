package nodes

import (
	"spiky/pkg/core"
	"spiky/pkg/kernels"
	"testing"
)

func TestNodeSpike(t *testing.T) {
	kernel := kernels.StdpKernel{}
	node := Node(&kernel)
	time := core.Time(0)
	node.SetSpike(time, true)

	if node.GetSpike(time) != true {
		t.Error("Node should have spiked")
	}

	if node.GetSpike(core.Time(14)) != false {
		t.Error("Node should not spike at this time")
	}
}

func TestNodeRate(t *testing.T) {
	kernel := kernels.StdpKernel{}
	node := Node(&kernel)

	node.SetSpike(core.Time(0), true)
	node.SetSpike(core.Time(10), true)
	node.SetSpike(core.Time(40), false)
	node.SetSpike(core.Time(50), true)
	node.SetSpike(core.Time(80), false)
	node.SetSpike(core.Time(100), true)

	spikeRate, err := node.GetSpikeRate(0, 100)
	if err != nil {
		t.Error(err)
	}
	if spikeRate != 0.04 {
		t.Errorf("Invalid spike rate %f instead of 0.04", spikeRate)
	}
}
