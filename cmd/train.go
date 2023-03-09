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
