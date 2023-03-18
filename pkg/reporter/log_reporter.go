package reporter

import (
	"fmt"
	"spiky/pkg/core"
)

type LogReporter struct {
	latestMetrics *map[string]float64
}

func (obs *LogReporter) OnStart(model core.IModel, dataset core.IDataset, iterations int) {
}

func (obs *LogReporter) OnUpdate(metrics *map[string]float64) {
	obs.latestMetrics = metrics
}

func (obs *LogReporter) OnStop() {
	for k, v := range *obs.latestMetrics {
		fmt.Println(k, v)
	}
}

func NewLogReporter(trainer *core.Trainer) *LogReporter {
	reporter := &LogReporter{}
	trainer.Subscribe(reporter)
	return reporter
}
