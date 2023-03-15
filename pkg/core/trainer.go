package core

import (
	"math"
	"spiky/pkg/utils"
	"time"

	"github.com/aclements/go-gg/generic/slice"
)

type Trainer struct {
	model     IModel
	dataset   IDataset
	observers []IObserver
	constants *utils.Constants

	metrics     map[string]float64
	waitingTime int
	ticker      *time.Ticker
}

func (t *Trainer) Subscribe(observer IObserver) {
	t.observers = append(t.observers, observer)
}

func (t *Trainer) notify(fn func(obs IObserver)) {
	for _, obs := range t.observers {
		fn(obs)
	}
}

func (trainer *Trainer) SpeedDown() {
	trainer.waitingTime = utils.ClampInt(int(math.Floor(float64(trainer.waitingTime)*0.9-1)), 1, 10000)
	trainer.ticker = time.NewTicker(time.Duration(trainer.waitingTime) * time.Millisecond)
	trainer.metrics["waiting time"] = float64(trainer.waitingTime)
}

func (trainer *Trainer) SpeedUp() {
	trainer.waitingTime = utils.ClampInt(int(math.Ceil(float64(trainer.waitingTime)*1.1+1)), 1, 10000)
	trainer.ticker = time.NewTicker(time.Duration(trainer.waitingTime) * time.Millisecond)
	trainer.metrics["waiting time"] = float64(trainer.waitingTime)
}

func (trainer *Trainer) Train(epochs float64) {
	idx := 0
	model := trainer.model
	dataset := trainer.dataset
	datasetSize := trainer.dataset.Len()
	iterations := int(float64(datasetSize) * epochs)
	queueSize := 1000
	errors := utils.NewBooleanQueue(queueSize)

	trainer.notify(func(obs IObserver) { obs.OnStart(model, dataset, &trainer.metrics, iterations) })

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

		trainer.metrics["loss"] = loss
		trainer.metrics["expected"] = float64(expectedClass)
		trainer.metrics["predicted"] = float64(predictedClass)
		trainer.metrics["training"] = (float64(idx) / float64(iterations)) * 100
		trainer.metrics["time to fit"] = float64(endTime.Sub(startTime).Microseconds())
		trainer.metrics["error rate"] = float64(errors.Count()) / float64(errors.Len())

		trainer.notify(func(obs IObserver) { obs.OnUpdate(idx) })

		<-trainer.ticker.C
	}
	trainer.notify(func(obs IObserver) { obs.OnStop() })
}

func NewTrainer(model IModel, dataset IDataset, csts *utils.Constants) *Trainer {
	app := &Trainer{
		model:       model,
		dataset:     dataset,
		constants:   csts,
		observers:   []IObserver{},
		waitingTime: 0,
		ticker:      time.NewTicker(1),
		metrics:     make(map[string]float64),
	}
	return app
}
