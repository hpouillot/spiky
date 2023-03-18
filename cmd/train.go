/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/core/codec"
	"spiky/pkg/data"
	"spiky/pkg/reporter"
	"spiky/pkg/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Quiet bool

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetLevel(logrus.ErrorLevel)

		dataset := data.NewMnist("./mnist")
		inputSize, outputSize := dataset.Shape()
		csts := utils.NewDefaultConstants()
		// model := buildHiddenLayer(inputSize, outputSize, csts)
		model := buildPerceptron(inputSize, outputSize, csts)
		trainer := core.NewTrainer(model, dataset, csts)
		if !Quiet {
			reporter.NewTrainingReporter(trainer, csts)
		} else {
			reporter.NewProgressBarReporter(trainer)
		}
		reporter.NewLogReporter(trainer)
		trainer.Start(5)
	},
}

func buildPerceptron(inputSize int, outputSize int, csts *utils.Constants) *core.Model {
	codec := codec.NewLatencyCodec(255, csts)
	input := core.NewLayer("Input", inputSize)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, output, csts)
	layers := []*core.Layer{
		input,
		output,
	}
	model := core.NewModel(codec, layers, csts)
	return model
}

func buildHiddenLayer(inputSize int, outputSize int, csts *utils.Constants) *core.Model {
	codec := codec.NewLatencyCodec(255, csts)
	input := core.NewLayer("Input", inputSize)
	hidden := core.NewLayer("Hidden", 50)
	output := core.NewLayer("Output", outputSize)
	core.DenseConnection(input, hidden, csts)
	core.DenseConnection(hidden, output, csts)
	layers := []*core.Layer{
		input,
		hidden,
		output,
	}
	model := core.NewModel(codec, layers, csts)
	return model
}

func init() {
	trainCmd.Flags().BoolVarP(&Quiet, "quiet", "q", false, "Start training in quiet mode")
	rootCmd.AddCommand(trainCmd)
}
