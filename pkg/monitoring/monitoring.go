package monitoring

import (
	"log"
	"spiky/pkg/core"

	"github.com/gdamore/tcell/v2"
)

type Monitor struct {
	layer    core.Layer
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

func (m *Monitor) DrawNodeLayout() {
	pointStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	m.layer.Visit(func(node core.Node, idx int) {
		position := node.GetPosition()
		x := int(position.X * float64(m.width))
		y := int(position.Y * float64(m.height))
		m.screen.SetContent(x, y, tcell.RuneBullet, nil, pointStyle)
	})
}

func (m *Monitor) DrawNodeSpikes(duration int) {
	pointStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)
	nodeSize := m.layer.Size()
	m.layer.Visit(func(node core.Node, idx int) {
		for _, time := range node.GetSpikeTimes(0, core.Time(duration)) {
			x := int((float64(time.ToInt()) / float64(duration)) * float64(m.width))
			y := int((float64(idx) / float64(nodeSize)) * float64(m.height))
			m.screen.SetContent(x, y, tcell.RuneBullet, nil, pointStyle)
		}
	})
}

func (m *Monitor) Render(duration int) {
	m.screen.Clear()
	m.DrawNodeSpikes(duration)
	m.screen.Show()
}

func NewMonitor(layer core.Layer) Monitor {
	return Monitor{
		layer:    layer,
		isClosed: false,
	}
}
