package core

import (
	"math"
	"spiky/pkg/utils"
	"time"

	"github.com/aclements/go-gg/generic/slice"
)

type Trainer struct {
	model     *Model
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

func (trainer *Trainer) Start(epochs int) {

	trainer.stopped = false
	model := trainer.model
	dataset := trainer.dataset
	datasetSize := trainer.dataset.Len()
	queueSize := 1000
	errors := utils.NewBooleanQueue(queueSize)

	trainer.notify(func(obs IObserver) { obs.OnStart(model, dataset) })
	var metrics map[string]float64 = make(map[string]float64)

	for i := 0; i < epochs; i++ {
		idx := 0
		trainer.notify(func(obs IObserver) { obs.OnEpochStart(datasetSize) })
		for sample := range trainer.dataset.Cycle(datasetSize) {
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

			metrics["1. epoch"] = float64(i)
			metrics["2. step"] = float64(idx)
			metrics["3. success rate"] = 1.0 - float64(errors.Count())/float64(errors.Len())
			metrics["4. expected"] = float64(expectedClass)
			metrics["5. predicted"] = float64(predictedClass)
			metrics["6. completion"] = (float64(idx) / float64(datasetSize)) * 100
			metrics["7. fit duration"] = float64(endTime.Sub(startTime).Microseconds())

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

func NewTrainer(model *Model, dataset IDataset, csts *utils.Constants) *Trainer {
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
