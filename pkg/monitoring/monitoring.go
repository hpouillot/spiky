package monitoring

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type monitor struct {
	isClosed bool
	width    int
	height   int
	screen   tcell.Screen
}

func (m *monitor) open() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	m.screen = s
	m.isClosed = false
	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	m.screen.SetStyle(defStyle)
	m.width, m.height = m.screen.Size()
}

func (m *monitor) listenToEventsAsync() {
	for {
		// Poll event
		ev := m.screen.PollEvent()
		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			m.screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				m.close()
				break
			}
		}
	}
}

func (m *monitor) IsClosed() bool {
	return m.isClosed
}

func (m *monitor) close() {
	maybePanic := recover()
	m.screen.Fini()
	m.isClosed = true
	if maybePanic != nil {
		panic(maybePanic)
	}
}
