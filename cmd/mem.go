/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// memCmd represents the mem command
var (
	convert string
	memCmd  = &cobra.Command{
		Use:   "mem",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			url := utils.BuildURL(secure, host, port, "mem", filter)
			data, err := utils.MakeRequest(url)

			if err != nil {
				log.Fatal(err)
			}

			mem := gjson.ParseBytes(data)
			icinga := types.NewIcinga(fmt.Sprintf("mem %s = %f", onValue, mem.Get(onValue).Float()), warning, critical)

			// TODO convert mit einbauen
			icinga.Evaluate(mem.Get(onValue).Float(),
				"mem usage has problems",
				"",
				"",
				"",
			)
			icinga.ParseToPerfData(mem)

			icinga.GenerateOutput()
			os.Exit(icinga.ExitCode)
		},
	}
)

func init() {
	rootCmd.AddCommand(memCmd)

	memCmd.Flags().StringVar(&onValue, "on", "available", "The value on which field the value is checked")

	// own warning an critical flags for predefined defaults
	memCmd.Flags().StringVar(&warning, "warning", "1", "The value on which field the value is checked")
	memCmd.Flags().StringVar(&critical, "critical", "0.5", "The value on which field the value is checked")

	memCmd.Flags().StringVar(&convert, "convert", "GB", "The unit in which the data is converted. Supported is KB, MB, GB, TB")
}
