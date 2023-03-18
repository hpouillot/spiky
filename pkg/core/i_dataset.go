package core

type IDataset interface {
	Len() int
	Get(idx int) Sample
}

type Sample struct {
	X []float64
	Y []float64
}
