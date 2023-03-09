package data

type Generator[T interface{}] chan T

type Sample struct {
	X []byte
	Y []byte
}

type IDataset interface {
	Iter() Generator[Sample]
	Cycle() Generator[Sample]
}

type Dataset struct {
	xSize int
	ySize int

	xSamples [][]byte
	ySamples [][]byte
}

func (d *Dataset) Iter() Generator[Sample] {
	n := len(d.xSamples)
	ch := make(Generator[Sample])

	go func() {
		for i := 0; i < n; i++ {
			ch <- Sample{
				X: d.xSamples[i],
				Y: d.ySamples[i],
			}
		}
		close(ch)
	}()

	return ch
}

func (d *Dataset) Cycle(iterations int) Generator[Sample] {
	sampleSize := len(d.xSamples)
	ch := make(Generator[Sample])

	go func() {
		for i := 0; i < iterations; i++ {
			ch <- Sample{
				X: d.xSamples[i%sampleSize],
				Y: d.ySamples[i%sampleSize],
			}
		}
		close(ch)
	}()

	return ch
}

func (d *Dataset) Shape() (int, int) {
	return d.xSize, d.ySize
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
		xSize:    1,
		ySize:    1,
		xSamples: xSamples,
		ySamples: ySamples,
	}
}
