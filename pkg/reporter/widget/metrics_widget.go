package widget

import (
	"fmt"
	"sort"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"golang.org/x/exp/maps"
)

type MetricsWidget struct {
	widgets.Table
	metrics map[string]float64
}

func (m *MetricsWidget) Draw(buf *termui.Buffer) {
	m.Rows = make([][]string, len(m.metrics)+1)
	var keys []string = maps.Keys(m.metrics)
	sort.Strings(keys)
	m.Rows[0] = []string{"Metric", "Value"}
	idx := 0
	for _, key := range keys {
		m.Rows[idx+1] = []string{key, fmt.Sprintf("%.2f", m.metrics[key])}
		idx++
	}
	m.Table.Draw(buf)
}

func (m *MetricsWidget) Set(key string, value float64) {
	m.metrics[key] = value
}

func (m *MetricsWidget) Get(key string) float64 {
	return m.metrics[key]
}

func NewMetricsWidget() *MetricsWidget {
	widget := &MetricsWidget{
		Table:   *widgets.NewTable(),
		metrics: map[string]float64{},
	}

	return widget
}
