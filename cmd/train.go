/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/training"
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

		app := training.NewTrainingApp(model, dataset, csts)
		defer app.Close()
		app.Open()
		app.Start(55000)
	},
}

func buildModel(inputSize int, outputSize int, csts *utils.Constants) core.Model {
	codec := core.NewLatencyCodec(csts)
	input := core.NewLayer("Input", inputSize)
	// hidden1 := core.NewLayer("Hidden 1", 100)
	// core.DenseConnection(input, hidden1, csts)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, output, csts)
	layers := []*core.Layer{
		input,
		// hidden1,
		output,
	}
	model := core.NewSampleModel(codec, layers, csts)
	return model
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
