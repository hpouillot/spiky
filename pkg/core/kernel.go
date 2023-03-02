package core

type Kernel interface {
	Compute(node Node, time Time, queue *Queue)
	Update(node Node, time Time)
}
