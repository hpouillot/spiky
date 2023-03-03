package core

type Model interface {
	Train(duration Time)
	Run(duration Time)
	Reset()
	Visit(func(node Node, idx int))
}
