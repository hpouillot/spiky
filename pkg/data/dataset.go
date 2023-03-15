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

func (d *Dataset) Iter(iteration int) chan core.Sample {
	ch := make(chan core.Sample)

	go func(iter int) {
		for i := 0; i < iter; i++ {
			ch <- d.get(i)
		}
		close(ch)
	}(iteration)

	return ch
}

func (d *Dataset) Cycle(iterations int) chan core.Sample {
	ch := make(chan core.Sample)

	go func(iter int) {
		for i := 0; i < iter; i++ {
			ch <- d.get(i % d.len)
		}
		close(ch)
	}(iterations)

	return ch
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
