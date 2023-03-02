package monitoring

import (
	"image"
	"log"
	"spiky/pkg/core"

	ui "github.com/gizak/termui/v3"
)

type Monitor struct {
	layer      core.Layer
	isClosed   bool
	termWidth  int
	termHeight int
}

func (m *Monitor) Create() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	termWidth, termHeight := ui.TerminalDimensions()
	m.termHeight = termHeight
	m.termWidth = termWidth
	go func() {
		for e := range ui.PollEvents() {
			switch e.ID {
			case "q", "<C-c>":
				m.isClosed = true
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				m.termHeight = payload.Height
				m.termWidth = payload.Width
				ui.Clear()
				m.Render(0)
			}
		}
	}()
}

func (m *Monitor) Close() {
	ui.Close()
}

func (m *Monitor) IsClosed() bool {
	return m.isClosed
}

func (m *Monitor) Render(step int) {
	c := ui.NewCanvas()
	c.SetRect(0, 0, m.termWidth, m.termHeight)
	m.layer.Visit(func(node core.Node) {
		position := node.GetPosition()
		x := int(position.X * float64(m.termWidth))
		y := int(position.Y * float64(m.termHeight))
		point := image.Pt(x, y)
		c.SetPoint(point, ui.ColorYellow)
	})
	ui.Render(c)
}

func NewMonitor(layer core.Layer) Monitor {
	return Monitor{
		layer:    layer,
		isClosed: false,
	}
}
