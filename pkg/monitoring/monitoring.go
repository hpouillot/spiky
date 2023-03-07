package monitoring

import (
	"log"
	"spiky/pkg/core"

	"github.com/gdamore/tcell/v2"
)

type Monitor struct {
	layer    core.Box[core.Neuron]
	isClosed bool
	width    int
	height   int
	screen   tcell.Screen
}

func (m *Monitor) Create() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	m.screen = s
	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	m.screen.SetStyle(defStyle)

	// Clear screen
	m.screen.Clear()
	m.width, m.height = m.screen.Size()
	go (func() {
		for {
			// Poll event
			ev := s.PollEvent()

			// Process event
			switch ev := ev.(type) {
			case *tcell.EventResize:
				s.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					m.isClosed = true
					return
				} else if ev.Key() == tcell.KeyCtrlL {
					s.Sync()
				} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
					s.Clear()
				}
			}
		}
	})()
}

func (m *Monitor) Close() {
	maybePanic := recover()
	m.screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func (m *Monitor) IsClosed() bool {
	return m.isClosed
}

func (m *Monitor) DrawNodeSpikes(duration float64) {
	pointStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	nodeSize := m.layer.Size()
	m.layer.Visit(func(idx int, node *core.Neuron) {
		for _, time := range (*node.GetSpikes())[0:int(duration)] {
			x := int((time / duration) * float64(m.width))
			y := int((float64(idx) / float64(nodeSize)) * float64(m.height))
			m.screen.SetContent(x, y, tcell.RuneBlock, nil, pointStyle)
		}
	})
}

func (m *Monitor) Render(duration float64) {
	m.screen.Clear()
	m.DrawNodeSpikes(duration)
	m.screen.Show()
}

func NewMonitor(layer core.Box[core.Neuron]) Monitor {
	monitor := Monitor{
		layer:    layer,
		isClosed: false,
	}
	return monitor
}
