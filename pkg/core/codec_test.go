package core

import (
	"testing"
)

func TestNewRateCodec(t *testing.T) {
	codec := NewRateCodec(3)
	spikes := codec.Encode(255)
	if len(spikes) < 5 {
		t.Error("Invalid spike count")
	}
}
