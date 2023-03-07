/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// source := data.Byte([]byte{
		// 	10,
		// 	250,
		// }) // Sized, Localized dataset ?

		// kernel := kernels.StdpKernel{
		// 	Threshold:      200.0,
		// 	Tho:            20,
		// 	LearningRate:   0.1,
		// 	MaxWeight:      200.0,
		// 	TraceTarget:    1,
		// 	RefractoryTime: 2.0,
		// 	MaxDelay:       2.0,
		// }

		// input := layers.Input(source, 1)
		// layer1 := layers.Layer(10, &kernel)
		// layer2 := layers.Layer(5, &kernel)

		// edges.Dense(input, layer1, 1.0, kernel.MaxWeight, kernel.MaxDelay)
		// edges.Dense(layer1, layer2, 1.0, kernel.MaxWeight, kernel.MaxDelay)

		// model := models.Model(input, layer1)
		// monitor := monitoring.NewMonitor(layer1)
		// monitor.Create()
		// defer monitor.Close()

		// iteration := 1000
		// runDuration := core.Time(100.0)
		// for k := 0; k < iteration; k++ {
		// 	model.Run(runDuration)
		// 	monitor.Render(runDuration)
		// 	model.Reset()
		// 	source.Next(true)
		// 	time.Sleep(1000 * time.Millisecond)
		// 	if monitor.IsClosed() {
		// 		break
		// 	}
		// }
		model := buildModel()
		fmt.Println(model)
	},
}

func buildModel() int {
	return 0
}

func init() {
	rootCmd.AddCommand(trainCmd)
}
