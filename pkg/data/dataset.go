package data

type Generator[T interface{}] chan T

type Sample struct {
	X []byte
	Y []byte
}

type Shape struct {
	X int
	Y int
}

type Dataset struct {
	len   int
	shape Shape
	get   func(idx int) Sample
}

func (d *Dataset) Iter() Generator[Sample] {
	ch := make(Generator[Sample])

	go func(iter int) {
		for i := 0; i < iter; i++ {
			ch <- d.get(i)
		}
		close(ch)
	}(d.len)

	return ch
}

func (d *Dataset) Cycle(iterations int) Generator[Sample] {
	ch := make(Generator[Sample])

	go func(iter int) {
		for i := 0; i < iter; i++ {
			ch <- d.get(i % d.len)
		}
		close(ch)
	}(iterations)

	return ch
}

func (d *Dataset) Len() int {
	return d.len
}

func (d *Dataset) Shape() (int, int) {
	return d.shape.X, d.shape.Y
}

func NewNumberDataset(XSamples []byte, YSamples []byte) *Dataset {
	if len(XSamples) != len(YSamples) {
		panic("Samples must have the same size")
	}
	xSamples := make([][]byte, len(XSamples))
	ySamples := make([][]byte, len(YSamples))
	for idx, X := range XSamples {
		xSamples[idx] = []byte{X}
	}
	for idx, Y := range YSamples {
		ySamples[idx] = []byte{Y}
	}
	return &Dataset{
		get: func(idx int) Sample {
			return Sample{
				X: []byte{XSamples[idx]},
				Y: []byte{YSamples[idx]},
			}
		},
		len: len(xSamples),
		shape: Shape{
			X: 1,
			Y: 1,
		},
	}
}
