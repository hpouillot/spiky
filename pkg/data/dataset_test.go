package data

import (
	"fmt"
	"testing"
)

func TestDataset(t *testing.T) {
	dataset := NewNumberDataset([]float64{150, 100}, []float64{24, 25})
	for sample := range dataset.Iter(2) {
		fmt.Print(sample)
	}
}
