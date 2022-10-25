/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// diskCmd represents the disk command
var (
	onValue string

	disksCmd = &cobra.Command{
		Use:   "disks",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			url := utils.BuildURL(secure, host, port, "disks", filter)
			data, err := utils.MakeRequest(url)

			if err != nil {
				log.Fatal(err)
			}

			disks := gjson.ParseBytes(data)
			icinga := types.NewIcinga("OK - All Disks are working", warning, critical)

			for _, value := range disks.Array() {
				icinga.Evaluate(value.Get(onValue).Float(),
					"Some disks have problems",
					fmt.Sprintf("disk '%s' fits in the range with value of %f", value.Get("usage.path").String(), value.Get(onValue).Float()),
					fmt.Sprintf("disk '%s' exceeds the limit of warning %f with value of %f", value.Get("usage.path").String(), icinga.Warning.Up, value.Get(onValue).Float()),
					fmt.Sprintf("disk '%s' exceeds the limit of critical %f with value of %f", value.Get("usage.path").String(), icinga.Critical.Up, value.Get(onValue).Float()),
				)

				icinga.AddPerfData(value.Get(onValue).Float(), value.Get("usage.path").String())
			}

			icinga.GenerateOutput()
			os.Exit(icinga.ExitCode)
		},
	}
)

func init() {
	rootCmd.AddCommand(disksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	disksCmd.Flags().StringVar(&onValue, "on", "usage.usedPercent", "The value on which field the value is checked")

	disksCmd.Flags().StringVar(&warning, "warning", "90", "The value on which field the value is checked")
	disksCmd.Flags().StringVar(&critical, "critical", "95", "The value on which field the value is checked")

}
