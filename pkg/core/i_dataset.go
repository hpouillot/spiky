package core

type IDataset interface {
	Len() int
	Cycle(size int) chan Sample
	Iter(size int) chan Sample
}

type Sample struct {
	X []byte
	Y []byte
}
