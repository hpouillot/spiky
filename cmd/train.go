/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"spiky/pkg/edges"
	"spiky/pkg/env"
	"spiky/pkg/layers"
	"spiky/pkg/models"

	"github.com/spf13/cobra"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		source := env.Text([]string{
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
		}) // Sized, Localized dataset ?
		target := env.Text([]string{
			"Y",
			"N",
			"Y",
			"N",
			"Y",
			"N",
		})

		input := layers.Input(source)
		output := layers.Output(target)

		edges.Dense(input, output, 0.5)

		model := models.Model(input, output)

		for k := 0; k < 5; k++ {
			model.Run(10000)
			source.Next()
			target.Next()
		}
		fmt.Println("train called")
	},
}

func init() {
	rootCmd.AddCommand(trainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
