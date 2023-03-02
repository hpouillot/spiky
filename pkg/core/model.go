package core

type Model interface {
	Train(duration int)
	Run(duration int)
	Reset()
	Visit(func(node Node, idx int))
}
