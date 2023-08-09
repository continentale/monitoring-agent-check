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

// modeDisksINodes represents the disks inodes command
var modeDisksINodes = &cobra.Command{
	Use:   "inodes",
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

		icinga := icinga.NewIcinga("All Disks has sufficient inodes", warning, critical)

		for _, value := range disks {
			icinga.Evaluate(value.Usage.InodesUsedPercent,
				"Some disks have not sufficient inodes",
				fmt.Sprintf("disk '%s' fits in the range with value of %f", value.Usage.Path, value.Usage.InodesUsedPercent),
				fmt.Sprintf("disk '%s' exceeds the limit of warning %f with value of %f", value.Usage.Path, icinga.Warning.Up, value.Usage.InodesUsedPercent),
				fmt.Sprintf("disk '%s' exceeds the limit of critical %f with value of %f", value.Usage.Path, icinga.Critical.Up, value.Usage.InodesUsedPercent),
			)
			icinga.AddPerfData(value.Usage.InodesUsedPercent, value.Usage.Path)
		}

		icinga.GenerateOutput(perfData)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	disksCmd.AddCommand(modeDisksINodes)
}
