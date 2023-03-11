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
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// modeMemUsageCmd represents the mem usage command
var modeMemUsageCmd = &cobra.Command{
	Use:   "used",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildURL(secure, host, port, "mem", filter)
		if verbose {
			log.Println("Used url:", url)
		}
		data, err := utils.MakeRequest(url)

		if err != nil {
			log.Fatal(err)
		}

		var mem mem.VirtualMemoryStat
		err = json.Unmarshal(data, &mem)
		if err != nil {
			log.Fatal(err)
		}

		icinga := icinga.NewIcinga(fmt.Sprintf("mem usage = %f", mem.UsedPercent), warning, critical)

		icinga.Evaluate(float64(mem.UsedPercent),
			"mem usage has problems",
			fmt.Sprintf("mem usage %f", mem.UsedPercent),
			fmt.Sprintf("mem usage %f", mem.UsedPercent),
			fmt.Sprintf("mem usage %f", mem.UsedPercent),
		)

		jsonData, err := json.Marshal(mem)
		if err != nil {
			log.Fatal("Cannot create perf Data", err)
		}
		icinga.ParseToPerfData(gjson.ParseBytes(jsonData))

		icinga.GenerateOutput(false)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	memCmd.AddCommand(modeMemUsageCmd)
}
