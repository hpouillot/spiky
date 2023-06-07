/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/factory"
	"spiky/pkg/reporter"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Quiet bool

var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetLevel(logrus.ErrorLevel)

		dataset := data.NewMnist("./mnist")
		inputSize, outputSize := dataset.Shape()
		config := core.NewDefaultConfig()
		world := core.NewWorld(config)
		model := factory.BuildSequential(world, []int{inputSize, outputSize})
		trainer := core.NewTrainer(model, dataset)
		if !Quiet {
			reporter.NewTrainingReporter(trainer, config)
		} else {
			reporter.NewProgressBarReporter(trainer)
			reporter.NewLogReporter(trainer)
		}
		trainer.Start(5)
	},
}

func init() {
	trainCmd.Flags().BoolVarP(&Quiet, "quiet", "q", false, "Start training in quiet mode")
	rootCmd.AddCommand(trainCmd)
}
