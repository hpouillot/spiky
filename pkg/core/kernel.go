package core

type Kernel interface {
	Compute(node *Node, time Time) bool
	Update(node *Node, time Time)
}
