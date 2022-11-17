/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/shirou/gopsutil/load"
	"github.com/spf13/cobra"
)

// modeMemUsageCmd represents the cpu load command
var modeCPULoad = &cobra.Command{
	Use:   "load",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildURL(secure, host, port, "load", filter)

		if verbose {
			log.Println("Used url:", url)
		}
		data, err := utils.MakeRequest(url)

		if err != nil {
			log.Fatal(err)
		}

		var load load.AvgStat
		err = json.Unmarshal(data, &load)
		if err != nil {
			log.Fatal(err)
		}
		icinga := types.NewIcinga("CPU usage: ", warning, critical)
		fmt.Println(icinga)

		fmt.Println(load)
	},
}

func init() {
	cpusCmd.AddCommand(modeCPULoad)
}
