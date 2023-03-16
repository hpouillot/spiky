package core

type IEncoder interface {
	Encode(value float64) []float64 // Schedule spikes for nodes
}

type IDecoder interface {
	Decode(spikes []float64) float64
}

type ICodec interface {
	IEncoder
	IDecoder
}
