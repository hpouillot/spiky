package reporter

import (
	"fmt"
	"spiky/pkg/core"
)

type LogReporter struct {
	latestMetrics *map[string]float64
}

func (obs *LogReporter) OnStart(model *core.Model, dataset core.IDataset) {
}

func (obs *LogReporter) OnStep(metrics *map[string]float64) {
	obs.latestMetrics = metrics
}

func (obs *LogReporter) OnEpochStart(iterations int) {

}

func (obs *LogReporter) OnEpochEnd() {
	obs.printMetrics()
}

func (obs *LogReporter) OnStop() {

}

func (obs *LogReporter) printMetrics() {
	fmt.Println("============================")
	for k, v := range *obs.latestMetrics {
		fmt.Println(k, v)
	}
	fmt.Println("============================")
}

func NewLogReporter(trainer *core.Trainer) *LogReporter {
	reporter := &LogReporter{}
	trainer.Subscribe(reporter)
	return reporter
}
