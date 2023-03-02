package core

type Dataset interface {
	Get(point Point, time Time) bool
	Next()
	Reset()
	Size() int
	HasNext() bool
}
