package data

import (
	"fmt"
	"testing"
)

func TestDataset(t *testing.T) {
	dataset := NewNumberDataset([]byte{150, 100}, []byte{24, 25})
	for sample := range dataset.Iter() {
		fmt.Print(sample)
	}
}
