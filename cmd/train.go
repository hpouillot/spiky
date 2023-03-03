/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"spiky/pkg/core"
	"spiky/pkg/data"
	"spiky/pkg/edges"
	"spiky/pkg/kernels"
	"spiky/pkg/layers"
	"spiky/pkg/models"
	"spiky/pkg/monitoring"
	"time"

	"github.com/spf13/cobra"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		source := data.Text([]string{
			"0",
			"1",
			"2",
			"3",
			"4",
		}) // Sized, Localized dataset ?

		kernel := kernels.StdpKernel{
			Threshold:      200.0,
			Tho:            20,
			LearningRate:   0.1,
			MaxWeight:      200.0,
			RefractoryTime: 1.0,
			TraceTarget:    1,
			MaxDelay:       2.0,
		}

		input := layers.Input(source, 1)
		layer1 := layers.Layer(10, &kernel)
		layer2 := layers.Layer(5, &kernel)
		// layer2 := layers.Layer(50, &kernel)

		edges.Dense(input, layer1, 1.0, kernel.MaxWeight, kernel.MaxDelay)
		// edges.Bidirectional(layer1, 1.0, kernel.MaxWeight, kernel.MaxDelay)
		edges.Dense(layer1, layer2, 1.0, kernel.MaxWeight, kernel.MaxDelay)
		// edges.Bidirectional(layer2, 1.0, kernel.MaxWeight, kernel.MaxDelay)

		model := models.Model(input, layer1)
		monitor := monitoring.NewMonitor(input)
		monitor.Create()

		defer monitor.Close()

		iteration := 1000
		runDuration := core.Time(100.0)
		for k := 0; k < iteration; k++ {
			source.Next(true)
			model.Run(runDuration)
			monitor.Render(runDuration)
			model.Reset()
			time.Sleep(1000 * time.Millisecond)
			if monitor.IsClosed() {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
