package monitoring

import (
	"spiky/pkg/core"
	"time"

	"github.com/gdamore/tcell/v2"
)

type SpikeMonitor struct {
	monitor
	layer      core.Box[core.Neuron]
	timeWindow int
}

func (m *SpikeMonitor) Open(ticker <-chan time.Time) {
	m.open()
	go m.listenToEventsAsync()
	go m.observe(ticker)
}

func (m *SpikeMonitor) observe(ticker <-chan time.Time) {
	pointStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	layerSize := m.layer.Size()
	for range ticker {
		m.screen.Clear()
		m.layer.Visit(func(idx int, n *core.Neuron) {
			for _, time := range *n.GetSpikes() {
				x := int((time / float64(m.timeWindow)) * float64(m.width))
				y := int((float64(idx) / float64(layerSize)) * float64(m.height))
				m.screen.SetContent(x, y, tcell.RuneBullet, nil, pointStyle)
			}
		})
		m.screen.Show()
	}
	m.close()
}

func NewSpikeMonitor(layer core.Box[core.Neuron], timeWindow int) *SpikeMonitor {
	return &SpikeMonitor{
		layer:      layer,
		timeWindow: timeWindow,
	}
}
