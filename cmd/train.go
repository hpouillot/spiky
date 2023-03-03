/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
			Threshold:    250.0,
			Tho:          50,
			LearningRate: 0.1,
			MaxWeight:    250.0,
		}

		input := layers.Input(source)
		layer1 := layers.Layer(100, &kernel)
		// layer2 := layers.Layer(10, &kernel)

		edges.Dense(input, layer1, 1.0)
		// edges.Dense(layer1, layer1, 0.1)

		model := models.Model(input, layer1)
		monitor := monitoring.NewMonitor(layer1)
		monitor.Create()

		defer monitor.Close()

		iteration := 300
		runDuration := 1000
		for k := 0; k < iteration; k++ {
			source.Next(true)
			model.Run(runDuration)
			monitor.Render(runDuration)
			model.Reset()
			time.Sleep(500 * time.Millisecond)
			if monitor.IsClosed() {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
