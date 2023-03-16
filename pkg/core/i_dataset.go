package core

type IDataset interface {
	Len() int
	Get(idx int) Sample
	Cycle(size int) chan Sample
	Iter(size int) chan Sample
}

type Sample struct {
	X []float64
	Y []float64
}
