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
	Short: "checks the load of the system",
	Long: `Checks the load endpoint of the agent. Evaluates the load5 metric each time the agent is running
	
	Examples:
	

	./monitoring-agent --host localhost cpus load
	can result in: OK - load usage is ok | 'load1'=0.030000;90.000000;95.000000;;'load5'=0.060000;90.000000;95.000000;;'load15'=0.020000;90.000000;95.000000;;

	./monitoring-agent --host localhost cpus load --warning 15 --critical 30
	./monitoring-agent --host localhost cpus load --perf=false
	`,
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
		icinga := types.NewIcinga("load usage is ok", warning, critical)
		icinga.InlineEvaluate(load.Load5,
			"Load value is high",
			fmt.Sprintf("load fits in the range with value of %f", load.Load5),
			fmt.Sprintf("load exceeds the limit of warning %f with value of %f", icinga.Warning.Up, load.Load5),
			fmt.Sprintf("load exceeds the limit of critical %f with value of %f", icinga.Critical.Up, load.Load5),
			verbose,
		)
		icinga.AddPerfData(load.Load1, "load1")
		icinga.AddPerfData(load.Load5, "load5")
		icinga.AddPerfData(load.Load15, "load15")
		icinga.GenerateOutput(perfData)
	},
}

func init() {
	cpusCmd.AddCommand(modeCPULoad)
}
