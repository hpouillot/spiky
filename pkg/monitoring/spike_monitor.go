package monitoring

import (
	"spiky/pkg/core"
	"time"
)

type SpikeMonitor struct {
	ticker *time.Ticker
	layer  core.Box[core.Neuron]
}

func (m *SpikeMonitor) Open() {
	go m.Refresh(m.ticker)
}

func (m *SpikeMonitor) Refresh(ticker *time.Ticker) {
	for range m.ticker.C {
		m.layer.Visit(func(idx int, n *core.Neuron) {

		})
	}
}

func (m *SpikeMonitor) Close() {
	m.ticker.Stop()
}

func NewSpikeMonitor(layer core.Box[core.Neuron]) *SpikeMonitor {
	return &SpikeMonitor{
		layer:  layer,
		ticker: time.NewTicker(30 * time.Millisecond),
	}
}
