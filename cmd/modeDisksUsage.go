/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/continentale/monitoring-agent-check/icinga"
	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
)

// modeDisksUsage represents the disks usage command
var modeDisksUsage = &cobra.Command{
	Use:   "usage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildFilterURL(secure, host, port, "disks", filter)
		if verbose {
			log.Println("Used url:", url)
		}
		data, err := utils.MakeRequest(url)

		if err != nil {
			log.Fatal(err)
		}

		var disks []types.Disks
		err = json.Unmarshal(data, &disks)
		if err != nil {
			log.Fatal(err)
		}

		icinga := icinga.NewIcinga("All Disks has sufficient space", warning, critical)
		for _, value := range disks {
			icinga.MultiEvaluate(value.Usage.UsedPercent,
				"Some disks have not sufficient disk space",
				fmt.Sprintf("disk '%s' fits in the range with value of %f", value.Usage.Path, value.Usage.UsedPercent),
				fmt.Sprintf("disk '%s' exceeds the threshold of warning '%s' with value of %f", value.Usage.Path, icinga.Warning.Raw, value.Usage.UsedPercent),
				fmt.Sprintf("disk '%s' exceeds the threshold of critical '%s' with value of %f", value.Usage.Path, icinga.Critical.Raw, value.Usage.UsedPercent),
			)
			icinga.AddPerfData(value.Usage.UsedPercent, value.Usage.Path)
		}
		icinga.GenerateOutput(perfData)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	disksCmd.AddCommand(modeDisksUsage)
}
