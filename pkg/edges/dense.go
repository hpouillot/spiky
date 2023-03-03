package edges

import (
	"spiky/pkg/core"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

func Dense(inputLayer core.Layer, outputLayer core.Layer, proba float64, maxWeight float64, maxDelay float64) {
	src := rand.New(rand.NewSource(1))
	distrib := distuv.Bernoulli{P: proba, Src: src}

	inputLayer.Visit(func(source core.Node, idx int) {
		outputLayer.Visit(func(target core.Node, idx int) {
			if distrib.Rand() == 1 {
				weight := rand.Float64() * maxWeight
				delay := rand.Float64() * maxDelay
				New(source, target, weight, delay)
			}
		})
	})
}

func Bidirectional(inputLayer core.Layer, proba float64, maxWeight float64, maxDelay float64) {
	src := rand.New(rand.NewSource(1))
	distrib := distuv.Bernoulli{P: proba, Src: src}

	inputLayer.Visit(func(source core.Node, idx1 int) {
		inputLayer.Visit(func(target core.Node, idx2 int) {
			if distrib.Rand() == 1 && idx1 != idx2 {
				weight := rand.Float64() * maxWeight
				delay := rand.Float64() * maxDelay
				New(source, target, weight, delay)
				New(target, source, -weight, delay)
			}
		})
	})
}
