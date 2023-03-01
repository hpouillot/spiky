/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"spiky/pkg/codec"
	"spiky/pkg/models"

	"github.com/spf13/cobra"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		stringCodec := codec.StringCodec{}
		m := models.New[string, string](stringCodec, stringCodec, 10)
		m.Run(100)
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
