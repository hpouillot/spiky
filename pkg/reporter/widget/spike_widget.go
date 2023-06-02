package widget

import (
	"image"
	"spiky/pkg/core"

	"github.com/gizak/termui/v3"
)

type SpikeWidget struct {
	termui.Block
	layer      *core.Layer
	timeWindow int
}

func (m *SpikeWidget) SetLayer(layer *core.Layer) {
	m.layer = layer
}

func (m *SpikeWidget) Draw(buf *termui.Buffer) {
	m.Title = m.layer.GetName() + " Activation"
	m.Block.Draw(buf)
	layerSize := m.layer.Size()
	m.layer.Visit(func(idx int, n *core.Neuron) {
		spikeTime := n.GetSpikeTime()
		if spikeTime != nil {
			x := int((*spikeTime / float64(m.timeWindow)) * float64(m.Dx()))
			y := int((float64(idx) / float64(layerSize)) * float64(m.Dy()))
			cell := termui.NewCell('•', termui.NewStyle(termui.ColorWhite))
			buf.SetCell(cell, image.Pt(m.Inner.Min.X+x, m.Inner.Min.Y+y))
		}
	})
}

func NewSpikeWidget(layer *core.Layer, timeWindow int) *SpikeWidget {
	widget := &SpikeWidget{
		Block:      *termui.NewBlock(),
		layer:      layer,
		timeWindow: timeWindow,
	}
	widget.TitleStyle = termui.NewStyle(termui.ColorYellow)
	return widget
}
