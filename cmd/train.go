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
			"5",
			"6",
			"7",
			"8",
			"9",
		}) // Sized, Localized dataset ?
		kernel := kernels.StdpKernel{
			Threshold: 250.0,
			Tho:       10,
		}

		input := layers.Input(source)
		layer := layers.Layer(500, &kernel)

		edges.Dense(input, layer, 1.0)

		model := models.Model(input, layer)
		monitor := monitoring.NewMonitor(layer)
		monitor.Create()

		defer monitor.Close()

		for k := 0; k < 10000; k++ {
			source.Next(true)
			model.Run(500)
			monitor.Render(k)
			model.Reset()
			// time.Sleep(200 * time.Millisecond)
			if monitor.IsClosed() {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
