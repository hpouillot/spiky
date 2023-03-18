package widget

import (
	"fmt"
	"sort"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/exp/maps"
)

type MetricsWidget struct {
	widgets.List
	metrics map[string]float64
}

func (m *MetricsWidget) Draw(buf *termui.Buffer) {
	m.Rows = make([]string, len(m.metrics))
	var keys []string = maps.Keys(m.metrics)
	sort.Strings(keys)
	idx := 0
	for _, key := range keys {
		m.Rows[idx] = fmt.Sprintf("%v %.2f", key, m.metrics[key])
		idx++
	}
	m.List.Draw(buf)
}

func (m *MetricsWidget) Set(key string, value float64) {
	m.metrics[key] = value
}

func (m *MetricsWidget) Get(key string) float64 {
	return m.metrics[key]
}

func NewMetricsWidget() *MetricsWidget {
	widget := &MetricsWidget{
		List:    *widgets.NewList(),
		metrics: map[string]float64{},
	}
	widget.Title = "Metrics"
	return widget
}
