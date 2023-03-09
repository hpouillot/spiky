package core

type Box[T interface{}] interface {
	Visit(func(idx int, value *T))
	Size() int
}
