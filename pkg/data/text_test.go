package data

import (
	"math"
	"spiky/pkg/core"
	"testing"
)

func testByte(dataset core.Dataset, byteValue byte, position core.Point, t *testing.T) {
	count := 0.0
	for _, value := range [10000]int{} {
		if dataset.Get(position, core.Time(value)) {
			count += 1.0
		}
	}
	decodedByte := int((count / 10000.0) * math.MaxUint8)
	if decodedByte > int(byteValue)+1 {
		t.Errorf("invalid encoding expected: %d, found: %d", int(byteValue), decodedByte)
	}
	if decodedByte < int(byteValue)-1 {
		t.Errorf("invalid encoding expected: %d, found: %d", int(byteValue), decodedByte)
	}
}

func TestText(t *testing.T) {
	dataset := Text([]string{
		"Hello",
		"Coucou",
	})
	if dataset.Shape()[0] != 6 {
		t.Error("Invalid size")
	}
	stringBytes := []byte("Hello")
	for idx, currentByte := range stringBytes {
		position := core.Point{
			X: float64(idx) / float64(dataset.Shape()[0]),
			Y: 0,
			Z: 0,
		}
		testByte(dataset, currentByte, position, t)
	}
}
