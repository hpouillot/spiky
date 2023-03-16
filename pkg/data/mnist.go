package data

import (
	"path"
	"spiky/pkg/core"

	loader "github.com/moverest/mnist"
)

type Mnist struct {
	Dataset
}

func NewMnist(folder string) *Dataset {
	images, err := loader.LoadImageFile(path.Join(folder, "train-images-idx3-ubyte.gz"))
	if err != nil {
		panic("Can't open image file")
	}
	labels, err := loader.LoadLabelFile(path.Join(folder, "train-labels-idx1-ubyte.gz"))
	if err != nil {
		panic("Can't open labels file")
	}
	mnist := &Dataset{
		shape: Shape{
			X: 28 * 28,
			Y: 10,
		},
		len: len(images),
		get: func(idx int) core.Sample {
			labelSlice := make([]float64, 10)
			labelSlice[labels[idx]] = 255
			imageSlice := make([]float64, 28*28)
			for idx, b := range images[idx] {
				imageSlice[idx] = float64(b)
			}
			return core.Sample{
				X: imageSlice,
				Y: labelSlice,
			}
		},
	}
	return mnist
}
