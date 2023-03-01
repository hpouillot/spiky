package models

import (
	"spiky/pkg/core"
)

func New[I interface{}, O interface{}](input core.Codec[I], output core.Codec[O], size int) core.Model {
	// input_size := input.Size()
	// output_size := output.Size()
	// hidden_size := size
	// total_size := input_size + output_size + hidden_size
	// all_nodes := make([]*core.Node, total_size)

	// input_nodes := make([]*core.Node, input_size)
	// for i := 0; i < input_size; i++ {
	// 	node := &core.Node{}
	// 	input_nodes[i] = node
	// 	all_nodes[i] = node
	// }

	// output_nodes := make([]*core.Node, output_size)
	// for i := 0; i < output_size; i++ {
	// 	node := &core.Node{}
	// 	output_nodes[i] = &core.Node{}
	// 	all_nodes[input_size+i] = node
	// }

	// input_interface := core.ModelInterface[I]{
	// 	Codec: input,
	// 	Nodes: input_nodes,
	// }

	// output_interface := core.ModelInterface[O]{
	// 	Codec: output,
	// 	Nodes: output_nodes,
	// }

	// kernel := kernels.StdpKernel{}

	// return core.Model[I, O]{
	// 	Input:  input_interface,
	// 	Output: output_interface,
	// 	Kernel: &kernel,
	// }
	return core.Model{}
}
