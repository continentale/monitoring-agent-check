/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/continentale/monitoring-agent-check/icinga"
	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
)

// modeMemUsageCmd represents the cpu usage command
var modeCPUUsage = &cobra.Command{
	Use:   "usage",
	Short: "checks the cpu usage of the system",
	Long: `checks the cpu endpoint of the agent. Evaluates the usage percent value
	
	./monitoring-agent --host localhost cpus usage`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildURL(secure, host, port, "cpus", filter)
		if filter[0] != "" {
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

		var cpus []types.CPUS
		err = json.Unmarshal(data, &cpus)
		if err != nil {
			log.Fatal(err)
		}
		icinga := icinga.NewIcinga("CPU usage: ", warning, critical)

		for _, value := range cpus {
			icinga.Evaluate(value.Usage,
				"cpu usage exceeds the threshold",
				fmt.Sprintf("cpu '%s' fits in the range with value of %f", value.TimeStat.CPU, value.Usage),
				fmt.Sprintf("cpu '%s' exceeds the limit of warning %f with value of %f", value.TimeStat.CPU, icinga.Warning.Up, value.Usage),
				fmt.Sprintf("cpu '%s' exceeds the limit of critical %f with value of %f", value.TimeStat.CPU, icinga.Critical.Up, value.Usage),
			)
			icinga.AddPerfData(value.Usage, value.TimeStat.CPU)
		}

		icinga.GenerateOutput(perfData)
	},
}

func init() {
	cpusCmd.AddCommand(modeCPUUsage)
}
