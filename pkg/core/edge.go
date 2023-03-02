package core

type Edge interface {
	GetTarget() Node
	GetSource() Node
	GetWeight() float64
	GetDelay() int
}
