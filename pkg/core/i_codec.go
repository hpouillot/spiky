package core

type IEncoder interface {
	Encode(layer *Layer, input []float64)
}

type IDecoder interface {
	Decode(layer *Layer) []float64
}

type IFitter interface {
	Fit(layer *Layer, output []float64)
}

type ICodec interface {
	IEncoder
	IDecoder
	IFitter
}
