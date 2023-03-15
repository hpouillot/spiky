package core

import (
	"spiky/pkg/utils"
	"time"

	"github.com/aclements/go-gg/generic/slice"
)

type Trainer struct {
	model     IModel
	dataset   IDataset
	observers []IObserver
	constants *utils.Constants
}

func (t *Trainer) Subscribe(observer IObserver) {
	t.observers = append(t.observers, observer)
}

func (t *Trainer) notify(fn func(obs IObserver)) {
	for _, obs := range t.observers {
		fn(obs)
	}
}

func (trainer *Trainer) Train(epochs float64) {
	idx := 0
	model := trainer.model
	dataset := trainer.dataset
	datasetSize := trainer.dataset.Len()
	iterations := int(float64(datasetSize) * epochs)
	queueSize := 1000
	errors := utils.NewBooleanQueue(queueSize)
	metrics := make(map[string]float64)
	trainer.notify(func(obs IObserver) { obs.OnStart(model, dataset, metrics, iterations) })
	for sample := range dataset.Cycle(iterations) {
		idx++
		startTime := time.Now()
		model.Reset()
		model.Encode(sample.X)
		model.Run()
		loss := model.Adjust(sample.Y)
		endTime := time.Now()
		predictions := model.Decode()
		predictedClass := slice.ArgMax(predictions)
		expectedClass := slice.ArgMax(sample.Y)
		errors.Push(predictedClass != expectedClass)

		metrics["loss"] = loss
		metrics["expected"] = float64(expectedClass)
		metrics["predicted"] = float64(predictedClass)
		metrics["training"] = (float64(idx) / float64(iterations)) * 100
		metrics["time to fit"] = float64(endTime.Sub(startTime).Microseconds())
		metrics["error rate"] = float64(errors.Count()) / float64(queueSize)
		trainer.notify(func(obs IObserver) { obs.OnUpdate(idx) })
	}
	trainer.notify(func(obs IObserver) { obs.OnStop() })
}

func NewTrainer(model IModel, dataset IDataset, csts *utils.Constants) *Trainer {
	app := &Trainer{
		model:     model,
		dataset:   dataset,
		constants: csts,
		observers: []IObserver{},
	}
	return app
}
