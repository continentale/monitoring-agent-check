/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

// swapCmd represents the swap command
var swapCmd = &cobra.Command{
	Use:   "swap",
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

		swapFreePercent := (mem.SwapFree / mem.SwapTotal) * 100

		icinga := types.NewIcinga(fmt.Sprintf("swap free: %d%%", swapFreePercent), warning, critical)

		icinga.InlineEvaluate(float64(mem.UsedPercent),
			"swap usage has exceed its limits",
			fmt.Sprintf("swap usage is ok: usage %d%%", swapFreePercent),
			fmt.Sprintf("swap usage has exceed its limits of %s with %d%%", warning, swapFreePercent),
			fmt.Sprintf("swap usage has exceed its limits of %s with %d%%", critical, swapFreePercent),
			verbose,
		)

		icinga.AddPerfData(float64(swapFreePercent), "swap")

		icinga.GenerateOutput(false)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	memCmd.AddCommand(swapCmd)
}
