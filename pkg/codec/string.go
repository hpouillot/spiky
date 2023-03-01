package codec

import "spiky/pkg/core"

type StringCodec struct {
	cat_codec CategoryCodec
}

func (s StringCodec) Size() int {
	return s.cat_codec.Size()
}

func (s StringCodec) Encode(data string, nodes []*core.Node) error {
	// codePoints := []rune(data)
	// return s.cat_codec.Encode(codePoints, nodes)
	return nil
}

func (s StringCodec) Decode(nodes []*core.Node) (string, error) {
	// bytes, err := s.b_codec.Decode(nodes)
	// if err != nil {
	// 	return "", err
	// }
	// return string(bytes), nil
	return "", nil
}

func NewStringCodec(length int, resolution int) *StringCodec {
	// codec := BytesCodec{
	// 	length:     length,
	// 	resolution: resolution,
	// }
	// return &StringCodec{
	// 	length:  length,
	// 	b_codec: codec,
	// }
	return new(StringCodec)
}
