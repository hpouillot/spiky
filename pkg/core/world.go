package core

import (
	"spiky/pkg/utils"

	"github.com/sirupsen/logrus"
)

type Process (func(world *World))

type World struct {
	time  float64
	Const *utils.Constants
	stack utils.TimeStack[Process]
}

func (w *World) GetTime() float64 {
	return w.time
}

func (w *World) setTime(time float64) {
	logrus.Infof("time %v", time)
	w.time = time
}

func (w *World) Schedule(time float64, process Process) {
	w.stack.Push(time, process)
}

func (w *World) Next(duration float64) bool {
	item := w.stack.Pop()
	if item == nil || w.time >= duration {
		return false
	}
	w.setTime(item.Time)
	item.Value(w)
	return true
}

func NewWorld(constants *utils.Constants) *World {
	return &World{
		time:  0.0,
		Const: constants,
		stack: *utils.NewTimeStack[Process](),
	}
}
