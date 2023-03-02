package core

type Layer interface {
	Reset()
	Size() int
	GetNode(idx int) Node
	Visit(func(Node))
}
