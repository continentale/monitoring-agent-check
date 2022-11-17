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
		url := utils.BuildURL(secure, host, port, "procs", filter)
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

		if filter == "" {
			filter = "procCount"
		}
		icinga := types.NewIcinga(fmt.Sprintf("proc count %d", len(procs)), warning, critical)
		icinga.AddPerfData(float64(len(procs)), filter)
		icinga.GenerateOutput()
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	procsCmd.AddCommand(modeProcsCount)
}
