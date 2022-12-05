/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
)

// readonlCmd represents the readonl command
var readonlCmd = &cobra.Command{
	Use:   "readonly",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildURL(secure, host, port, "disks", filter)
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

		icinga := types.NewIcinga("No readonly filesystem detected", "1", "1")

		icinga.Evaluate(
			float64(len(disks)),
			"Found readonly filesystems",
			"",
			"",
			"",
			verbose,
		)

		for _, value := range disks {
			if utils.ArrayContains("ro", value.Details.Opts) {
				icinga.LongPluginOutput += "Readonly disks: " + value.Details.Mountpoint + "\n"
			}
		}

		icinga.GenerateOutput(perfData)
		os.Exit(icinga.ExitCode)
	},
}

func init() {
	disksCmd.AddCommand(readonlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readonlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readonlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
