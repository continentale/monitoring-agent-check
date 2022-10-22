/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/continentale/monitoring-agent-check/types"
	"github.com/continentale/monitoring-agent-check/utils"
	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.BuildURL(secure, host, port, "disk", "")
		data, err := utils.MakeRequest(url)

		if err != nil {
			fmt.Println(err)
		}

		var disks []types.Disks
		err = json.Unmarshal(data, &disks)

		if err != nil {
			fmt.Println("error while parsing data:", err)
		}

		fmt.Println(disks)

	},
}

func init() {
	rootCmd.AddCommand(diskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
