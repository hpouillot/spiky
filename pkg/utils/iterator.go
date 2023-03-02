package utils

type Iterator[T interface{}] interface {
	HasNext() bool
	GetNext() T
}
