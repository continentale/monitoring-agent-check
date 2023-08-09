/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildFilterURL(secure, host, port, "time", filter)

		if verbose {
			log.Println("Used url:", url)
		}
		data, err := utils.MakeRequest(url)

		if err != nil {
			log.Fatal(err)
		}

		var t types.TimeSync
		err = json.Unmarshal(data, &t)
		if err != nil {
			log.Fatal(err)
		}

		timeDiff := math.Abs(float64(t.Timestamp) - float64(time.Now().Unix()))

		icinga := types.NewIcinga(fmt.Sprintf("Time ok - timestamp: %s time diff: %f", t.Formatted, timeDiff), warning, critical)

		icinga.InlineEvaluate(timeDiff, "Time ok - timestamp "+t.Formatted,
			fmt.Sprintf("Time ok - timestamp %s time diff %f", t.Formatted, timeDiff),
			fmt.Sprintf("Time WARNING - timestamp %s time diff %f", t.Formatted, timeDiff),
			fmt.Sprintf("Time CRITICAL - timestamp %s time diff %f", t.Formatted, timeDiff),
			verbose)
		icinga.GenerateOutput(perfData)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	RootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
