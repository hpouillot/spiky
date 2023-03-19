package widget

import (
	"fmt"
	"spiky/pkg/core"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type LayersWidget struct {
	*widgets.List
	layers []*core.Layer
}

func (m *LayersWidget) Draw(buf *termui.Buffer) {
	m.Rows = make([]string, len(m.layers))
	for idx, layer := range m.layers {
		m.Rows[idx] = fmt.Sprintf("[%v] %s (%d)", idx, layer.GetName(), layer.Size())
	}
	m.List.Draw(buf)
}

func NewLayersWidget(layers []*core.Layer) *LayersWidget {
	listWidget := widgets.NewList()
	listWidget.Title = "Layers"
	listWidget.TextStyle = termui.NewStyle(termui.ColorWhite)
	listWidget.WrapText = false
	listWidget.SelectedRowStyle = termui.NewStyle(termui.ColorCyan)

	widget := &LayersWidget{
		List:   listWidget,
		layers: layers,
	}
	widget.TitleStyle = termui.NewStyle(termui.ColorYellow)
	return widget
}
