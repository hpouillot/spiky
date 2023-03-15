package observer

import (
	"spiky/pkg/core"

	"github.com/schollz/progressbar/v3"
)

type ProgressBarObserver struct {
	bar *progressbar.ProgressBar
}

func (pb *ProgressBarObserver) OnStart(model core.IModel, dataset core.IDataset, metrics map[string]float64, iterations int) {
	pb.bar = progressbar.Default(int64(iterations))
}

func (pb *ProgressBarObserver) OnUpdate(idx int) {
	pb.bar.Add(1)
}

func (pb *ProgressBarObserver) OnStop() {
	pb.bar.Close()
}

func NewProgressBarObserver() *ProgressBarObserver {
	return &ProgressBarObserver{}
}
