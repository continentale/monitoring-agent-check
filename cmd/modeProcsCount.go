/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/continentale/monitoring-agent-check/icinga"
	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
)

// modeProcsCount represents the procs count command
var modeProcsCount = &cobra.Command{
	Use:   "count",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildFilterURL(secure, host, port, "procs", filter)
		if verbose {
			log.Println("Used url:", url)
		}
		data, err := utils.MakeRequest(url)

		if err != nil {
			log.Fatal(err)
		}

		var procs []types.Procs
		err = json.Unmarshal(data, &procs)
		if err != nil {
			log.Fatal(err)
		}

		icinga := icinga.NewIcinga(fmt.Sprintf("proc count %d for filter %s", len(procs), filter), warning, critical)
		icinga.AddPerfData(float64(len(procs)), fmt.Sprintf("%v", filter))
		icinga.Evaluate(
			float64(len(procs)),
			"proc count does not match the threshold",
			fmt.Sprintf("proc count match. Found: %d", len(procs)),
			fmt.Sprintf("proc count is warning. Found %d. Does not fit thresholds for filter %s", len(procs), strings.Join(filter, ",")),
			fmt.Sprintf("proc count is critical. Found %d. Does not fit thresholds for filter %s", len(procs), strings.Join(filter, ",")),
		)
		icinga.GenerateOutput(false)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	procsCmd.AddCommand(modeProcsCount)
}
