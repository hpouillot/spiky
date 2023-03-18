package reporter

import (
	"fmt"
	"spiky/pkg/core"

	"github.com/schollz/progressbar/v3"
)

type ProgressBarReporter struct {
	bar *progressbar.ProgressBar
}

func (pb *ProgressBarReporter) OnStart(model core.IModel, dataset core.IDataset, iterations int) {
	pb.bar = progressbar.Default(int64(iterations))
}

func (pb *ProgressBarReporter) OnUpdate(metrics *map[string]float64) {
	pb.bar.Add(1)
	pb.bar.Describe(fmt.Sprintf("%v", (*metrics)["1. success rate"]))
}

func (pb *ProgressBarReporter) OnStop() {
	pb.bar.Close()
}

func NewProgressBarReporter(trainer *core.Trainer) *ProgressBarReporter {
	reporter := &ProgressBarReporter{}
	trainer.Subscribe(reporter)
	return reporter
}
