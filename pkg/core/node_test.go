package core

import (
	"testing"
)

func TestNodeSpike(t *testing.T) {
	node := NewBaseNode()
	time := Time(0)
	node.SetSpike(time, true)

	if node.GetSpike(time) != true {
		t.Error("Node should have spiked")
	}

	if node.GetSpike(Time(14)) != false {
		t.Error("Node should not spike at this time")
	}
}

func TestNodeRate(t *testing.T) {
	node := NewBaseNode()

	node.SetSpike(Time(0), true)
	node.SetSpike(Time(10), true)
	node.SetSpike(Time(40), false)
	node.SetSpike(Time(50), true)
	node.SetSpike(Time(80), false)
	node.SetSpike(Time(100), true)

	spikeRate, err := node.GetSpikeRate(0, 100)
	if err != nil {
		t.Error(err)
	}
	if spikeRate != 0.04 {
		t.Errorf("Invalid spike rate %f instead of 0.04", spikeRate)
	}
}
