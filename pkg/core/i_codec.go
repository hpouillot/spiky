package core

type IEncoder interface {
	Encode(value byte) []float64 // Schedule spikes for nodes
}

type IDecoder interface {
	Decode(spikes []float64) byte
}

type ICodec interface {
	IEncoder
	IDecoder
}
