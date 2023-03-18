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

	stopped     bool
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
}

func (trainer *Trainer) GetWaitingTime() int {
	return trainer.waitingTime
}

func (trainer *Trainer) SpeedUp() {
	trainer.waitingTime = utils.ClampInt(int(math.Ceil(float64(trainer.waitingTime)*1.1+1)), 1, 10000)
	trainer.ticker = time.NewTicker(time.Duration(trainer.waitingTime) * time.Millisecond)
}

func (trainer *Trainer) Start(epochs float64) {
	idx := 0
	trainer.stopped = false
	model := trainer.model
	dataset := trainer.dataset
	datasetSize := trainer.dataset.Len()
	iterations := int(float64(datasetSize) * epochs)
	queueSize := 1000
	errors := utils.NewBooleanQueue(queueSize)

	trainer.notify(func(obs IObserver) { obs.OnStart(model, dataset, iterations) })
	var metrics map[string]float64 = make(map[string]float64)

	for sample := range trainer.dataset.Cycle(iterations) {
		idx++
		startTime := time.Now()
		model.Reset()
		model.Encode(sample.X)
		model.Run()
		model.Adjust(sample.Y)
		endTime := time.Now()
		predictions := model.Decode()
		predictedClass := slice.ArgMax(predictions)
		expectedClass := slice.ArgMax(sample.Y)
		errors.Push(predictedClass != expectedClass)

		metrics["0. step"] = float64(idx)
		metrics["1. success rate"] = 1.0 - float64(errors.Count())/float64(errors.Len())
		metrics["3. expected"] = float64(expectedClass)
		metrics["4. predicted"] = float64(predictedClass)
		metrics["5. completion"] = (float64(idx) / float64(iterations)) * 100
		metrics["6. fit duration"] = float64(endTime.Sub(startTime).Microseconds())

		trainer.notify(func(obs IObserver) { obs.OnUpdate(&metrics) })

		if idx >= iterations || trainer.stopped {
			break
		} else {
			<-trainer.ticker.C
		}
	}
	trainer.notify(func(obs IObserver) { obs.OnStop() })
}

func (trainer *Trainer) Stop() {
	trainer.stopped = true
}

func NewTrainer(model IModel, dataset IDataset, csts *utils.Constants) *Trainer {
	app := &Trainer{
		model:       model,
		dataset:     dataset,
		constants:   csts,
		observers:   []IObserver{},
		waitingTime: 1,
		ticker:      time.NewTicker(1),
	}
	return app
}
