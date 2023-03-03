package core

type Edge interface {
	GetTarget() Node
	GetSource() Node
	GetDelay() float64
	GetWeight() float64
	SetWeight(deltaW float64) float64
}
