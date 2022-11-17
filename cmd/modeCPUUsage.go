/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

// modeMemUsageCmd represents the cpu usage command
var modeCPUUsage = &cobra.Command{
	Use:   "usage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildURL(secure, host, port, "cpus", filter)
		if filter != "" {
			url += "&perCPU=" + strconv.FormatBool(perCPU)
		} else {
			url += "?perCPU=" + strconv.FormatBool(perCPU)
		}
		if verbose {
			log.Println("Used url:", url)
		}
		data, err := utils.MakeRequest(url)

		if err != nil {
			log.Fatal(err)
		}

		var cpus []cpu.TimesStat
		err = json.Unmarshal(data, &cpus)
		if err != nil {
			log.Fatal(err)
		}
		icinga := types.NewIcinga("CPU usage: ", warning, critical)
		fmt.Println(icinga)

		fmt.Println(cpus)
	},
}

func init() {
	cpusCmd.AddCommand(modeCPUUsage)
}
