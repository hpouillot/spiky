package core

import (
	"spiky/pkg/utils"
)

type Process (func(world *World))

type World struct {
	time  float64
	Const *utils.Constants
	stack *utils.TimeStack[Process]

	dirtyNeurons map[string]*Neuron
}

func (w *World) GetTime() float64 {
	return w.time
}

func (w *World) markDirty(n *Neuron) {
	w.dirtyNeurons[n.id] = n
}

func (w *World) setTime(time float64) {
	w.time = time
}

func (w *World) Schedule(time float64, process Process) {
	w.stack.Push(time, process)
}

func (w *World) Next() bool {
	item := w.stack.Pop()
	if item == nil || w.time >= w.Const.MaxTime {
		return false
	}
	w.setTime(item.Time)
	item.Value(w)
	return true
}

func (w *World) Reset() {
	for _, n := range w.dirtyNeurons {
		n.Reset()
	}
	w.dirtyNeurons = make(map[string]*Neuron)
	w.time = 0.0
	w.stack.Reset()
}

func NewWorld(constants *utils.Constants) *World {
	return &World{
		time:         0.0,
		Const:        constants,
		stack:        utils.NewTimeStack[Process](),
		dirtyNeurons: make(map[string]*Neuron),
	}
}
