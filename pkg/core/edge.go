package core

// type Edge struct {
// 	source *Node
// 	target *Node
// 	weight float32
// 	delay  int
// }

type Edge interface {
	GetTarget() Node
	GetSource() Node
	GetWeight() float32
	GetDelay() int
}
