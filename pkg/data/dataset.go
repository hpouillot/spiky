package data

import "spiky/pkg/core"

type Shape struct {
	X int
	Y int
}

type Dataset struct {
	len   int
	shape Shape
	get   func(idx int) core.Sample
}

func (d *Dataset) Get(idx int) core.Sample {
	return d.get(idx % d.len)
}

func (d *Dataset) Len() int {
	return d.len
}

func (d *Dataset) Shape() (int, int) {
	return d.shape.X, d.shape.Y
}
