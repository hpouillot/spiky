package core

type Edge interface {
	GetTarget() Node
	GetSource() Node
	GetDelay() float64
	GetWeight() float64
	UpdateWeight(deltaW float64) float64
}
