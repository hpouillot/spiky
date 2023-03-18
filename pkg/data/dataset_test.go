package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataset(t *testing.T) {
	dataset := NewNumberDataset([]float64{150, 100}, []float64{24, 25})
	sample1 := dataset.Get(0)
	sample2 := dataset.Get(1)
	assert.Equal(t, sample1.X[0], 150.0)
	assert.Equal(t, sample1.Y[0], 24.0)

	assert.Equal(t, sample2.X[0], 100.0)
	assert.Equal(t, sample2.Y[0], 25.0)
}
