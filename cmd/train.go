/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetLevel(logrus.ErrorLevel)
		dataset := data.NewNumberDataset([]byte{100, 100, 100, 100}, []byte{255, 255, 255, 255})
		inputSize, outputSize := dataset.Shape()
		model := buildModel(inputSize, outputSize)
		for sample := range dataset.Cycle(1000000) {
			model.Fit(sample.X, sample.Y)
		}
	},
}

func buildModel(inputSize int, outputSize int) core.Model {
	csts := utils.NewDefaultConstants()
	codec := core.NewLatencyCodec(csts)
	input := core.NewLayer(inputSize)
	output := core.NewLayer(outputSize)
	core.DenseConnection(input, output, csts)
	model := core.NewSampleModel(codec, input, output, csts)
	return model
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
