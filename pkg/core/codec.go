package core

type Encoder[T interface{}] interface {
	Encode(data T, nodes []*Node) error
	Size() int
}

type Decoder[T interface{}] interface {
	Decode(nodes []*Node) (T, error)
	Size() int
}

type Codec[T interface{}] interface {
	Encoder[T]
	Decoder[T]
}
