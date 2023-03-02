/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// source := data.Text([]string{
		// 	"1",
		// 	"2",
		// 	"3",
		// 	"4",
		// 	"5",
		// 	"6",
		// }) // Sized, Localized dataset ?
		// target := data.Text([]string{
		// 	"Y",
		// 	"N",
		// 	"Y",
		// 	"N",
		// 	"Y",
		// 	"N",
		// })

		// input := layers.Input(source)
		// output := layers.Output(target)

		// edges.Dense(input, output, 0.5)

		// model := models.Model(input, output)

		// for k := 0; k < 5; k++ {
		// 	model.Run(10000)
		// 	source.Next()
		// 	target.Next()
		// }

		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()

		p := widgets.NewParagraph()
		p.Text = "Hello World!"
		p.SetRect(0, 0, 25, 5)

		ui.Render(p)

		for e := range ui.PollEvents() {
			if e.Type == ui.KeyboardEvent {
				break
			}
		}
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
