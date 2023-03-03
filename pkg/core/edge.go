package core

type Edge interface {
	GetTarget() Node
	GetSource() Node
	GetDelay() int
	GetWeight() float64
	UpdateWeight(deltaW float64) float64
}
