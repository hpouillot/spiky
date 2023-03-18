package reporter

import (
	"fmt"
	"spiky/pkg/core"

	"github.com/schollz/progressbar/v3"
)

type ProgressBarReporter struct {
	bar *progressbar.ProgressBar
}

func (pb *ProgressBarReporter) OnStart(model *core.Model, dataset core.IDataset) {
}

func (pb *ProgressBarReporter) OnStep(metrics *map[string]float64) {
	pb.bar.Add(1)
	pb.bar.Describe(fmt.Sprintf("%v", (*metrics)["3. success rate"]))
}

func (pb *ProgressBarReporter) OnStop() {
	pb.bar.Close()
}

func (pb *ProgressBarReporter) OnEpochStart(iterations int) {
	pb.bar = progressbar.Default(int64(iterations))
}

func (pb *ProgressBarReporter) OnEpochEnd() {

}

func NewProgressBarReporter(trainer *core.Trainer) *ProgressBarReporter {
	reporter := &ProgressBarReporter{}
	trainer.Subscribe(reporter)
	return reporter
}
