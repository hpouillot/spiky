package core

type MockedCodec struct {
	encodeCallCount int
	decodeCallCount int
	sizeCallCount   int
}

func (mc *MockedCodec) Encode(data string, nodes []*Node) error {
	mc.encodeCallCount++
	return nil
}

func (mc *MockedCodec) Decode(nodes []*Node) (string, error) {
	mc.decodeCallCount++
	return "mocked", nil
}

func (mc *MockedCodec) Size() int {
	mc.sizeCallCount++
	return 10
}
