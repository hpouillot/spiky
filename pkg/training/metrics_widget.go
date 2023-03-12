package training

import (
	"fmt"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type MetricsWidget struct {
	widgets.List
	metrics *map[string]float64
}

func (m *MetricsWidget) Draw(buf *termui.Buffer) {
	m.Rows = make([]string, len(*m.metrics))
	idx := 0
	for k, v := range *m.metrics {
		m.Rows[idx] = fmt.Sprintf("%v %v", k, v)
		idx++
	}
	m.List.Draw(buf)
}

func NewMetricsWidget(metrics *map[string]float64) *MetricsWidget {
	return &MetricsWidget{
		List:    *widgets.NewList(),
		metrics: metrics,
	}
}
