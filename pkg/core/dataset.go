package core

type Dataset interface {
	Get(point Point, time Time) bool
	Next()
	Reset()
	Shape() [3]int
	HasNext() bool
}
