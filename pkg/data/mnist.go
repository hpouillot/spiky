package data

import (
	"path"

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
			Y: 1,
		},
		len: len(images),
		get: func(idx int) Sample {
			return Sample{
				X: []byte((*images[idx])[:]),
				Y: []byte{byte(labels[idx])},
			}
		},
	}
	return mnist
}
