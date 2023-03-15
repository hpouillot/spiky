/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/core/codec"
	"spiky/pkg/data"
	"spiky/pkg/observer"
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

		dataset := data.NewMnist("./mnist")
		inputSize, outputSize := dataset.Shape()
		csts := utils.NewDefaultConstants()
		model := buildModel(inputSize, outputSize, csts)
		trainer := core.NewTrainer(model, dataset, csts)
		observer.NewTrainingObserver(trainer, csts)
		trainer.Start(10000)
	},
}

func buildModel(inputSize int, outputSize int, csts *utils.Constants) core.IModel {
	codec := codec.NewLatencyCodec(csts)
	input := core.NewLayer("Input", inputSize)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, output, csts)
	layers := []*core.Layer{
		input,
		output,
	}
	model := core.NewSampleModel(codec, layers, csts)
	return model
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
