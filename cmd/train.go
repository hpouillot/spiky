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
			"Hello this is a sentence",
			"Coucou momo",
		}) // Sized, Localized dataset ?

		kernel := kernels.StdpKernel{
			Threshold:      200.0,
			Tho:            20,
			LearningRate:   0.0001,
			MaxWeight:      200.0,
			RefractoryTime: 1.0,
			TraceTarget:    0.5,
		}

		input := layers.Input(source)
		layer1 := layers.Layer(50, &kernel)
		layer2 := layers.Layer(50, &kernel)

		edges.Dense(input, layer1, 1.0)
		// edges.Dense(layer1, layer1, 0.2)
		edges.Dense(layer1, layer2, 1.0)

		model := models.Model(input, layer2)
		monitor := monitoring.NewMonitor(layer2)
		monitor.Create()

		defer monitor.Close()

		iteration := 1000
		runDuration := core.Time(20.0)
		for k := 0; k < iteration; k++ {
			source.Next(true)
			model.Run(runDuration)
			monitor.Render(runDuration)
			model.Reset()
			time.Sleep(200 * time.Millisecond)
			if monitor.IsClosed() {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
