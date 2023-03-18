package core

type IEncoder interface {
	Encode(value *float64) *float64
}

type IDecoder interface {
	Decode(time *float64) *float64
}

type ICodec interface {
	IEncoder
	IDecoder
}
