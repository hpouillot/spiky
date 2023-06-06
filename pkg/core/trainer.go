package core

import (
	"math"
	"spiky/pkg/utils"
	"time"
)

type Trainer struct {
	model     *Model
	dataset   IDataset
	observers []IObserver

	stopped bool
	speed   float64
	ticker  *time.Ticker
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
	trainer.speed = utils.ClampFloat(trainer.speed*0.9, 0.0001, 1.0)
	trainer.updateTicker()
}

func (trainer *Trainer) SpeedUp() {
	trainer.speed = utils.ClampFloat(trainer.speed*1.1, 0.0001, 1.0)
	trainer.updateTicker()
}

func (trainer *Trainer) updateTicker() {
	delay := time.Duration(1001-math.Floor(trainer.speed*1000)) * time.Millisecond
	trainer.ticker = time.NewTicker(delay)
}

func (trainer *Trainer) GetSpeed() float64 {
	return trainer.speed
}

func (trainer *Trainer) Start(epochs int) {

	trainer.stopped = false
	model := trainer.model
	dataset := trainer.dataset
	datasetSize := trainer.dataset.Len()
	queueSize := 1000
	errors := utils.NewBooleanQueue(queueSize)

	trainer.notify(func(obs IObserver) { obs.OnStart(model, dataset) })
	var metrics map[string]float64 = make(map[string]float64)
	var sample Sample
	for i := 0; i < epochs; i++ {
		idx := 0
		trainer.notify(func(obs IObserver) { obs.OnEpochStart(datasetSize) })
		for j := 0; j < datasetSize; j++ {
			sample = dataset.Get(j)
			idx++
			model.Reset()
			model.Encode(sample.X)
			model.Run()
			model.Adjust(sample.Y)

			predictions := model.Decode()
			predictedClass := utils.ArgMax(predictions)
			expectedClass := utils.ArgMax(sample.Y)
			errors.Push(predictedClass != expectedClass)

			metrics["1. success %"] = (1.0 - float64(errors.Count())/float64(errors.Len())) * 100
			metrics["2. epoch"] = float64(i)
			metrics["3. completion %"] = (float64(idx) / float64(datasetSize)) * 100
			metrics["4. expected"] = float64(expectedClass)
			metrics["5. predicted"] = float64(predictedClass)

			trainer.notify(func(obs IObserver) { obs.OnStep(&metrics) })

			if trainer.stopped {
				break
			} else {
				<-trainer.ticker.C
			}
		}
		trainer.notify(func(obs IObserver) { obs.OnEpochEnd() })
		model.World.Const.LearningRate = model.World.Const.LearningRate * 0.1
	}
	trainer.notify(func(obs IObserver) { obs.OnStop() })
}

func (trainer *Trainer) Stop() {
	trainer.stopped = true
}

func NewTrainer(model *Model, dataset IDataset) *Trainer {
	app := &Trainer{
		model:     model,
		dataset:   dataset,
		observers: []IObserver{},
		speed:     1.0,
		ticker:    time.NewTicker(1),
	}
	return app
}
