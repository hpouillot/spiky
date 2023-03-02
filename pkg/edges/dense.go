package edges

import (
	"spiky/pkg/core"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

func Dense(inputLayer core.Layer, outputLayer core.Layer, proba float64) {
	src := rand.New(rand.NewSource(1))
	distrib := distuv.Bernoulli{P: proba, Src: src}

	inputLayer.Visit(func(source core.Node, idx int) {
		outputLayer.Visit(func(target core.Node, idx int) {
			if distrib.Rand() == 1 {
				source.Connect(target)
			}
		})
	})
}
