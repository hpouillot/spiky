/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/monitoring"
	"spiky/pkg/utils"
	"time"

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
		// dataset := data.NewNumberDataset([]byte{100, 3, 255, 15, 26}, []byte{255, 43, 12, 54, 23})
		inputSize, outputSize := dataset.Shape()
		csts := utils.NewDefaultConstants()
		model := buildModel(inputSize, outputSize, csts)
		monitor := monitoring.NewSpikeMonitor(model.GetInput(), int(csts.MaxTime))
		timeChannel := make(chan time.Time)
		monitor.Open(timeChannel)
		for sample := range dataset.Cycle(1000) {
			model.Fit(sample.X, sample.Y)
			timeChannel <- time.Now()
			time.Sleep(500 * time.Millisecond)
			if monitor.IsClosed() {
				break
			}
		}
		close(timeChannel)
	},
}

func buildModel(inputSize int, outputSize int, csts *utils.Constants) core.Model {
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
