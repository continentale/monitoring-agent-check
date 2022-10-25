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

// procCmd represents the proc command
var (
	count    bool
	procsCmd = &cobra.Command{
		Use:   "procs",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			url := utils.BuildURL(secure, host, port, "procs", filter)
			data, err := utils.MakeRequest(url)

			if err != nil {
				fmt.Println(err)
			}

			var procs []types.Procs
			err = json.Unmarshal(data, &procs)

			if err != nil {
				fmt.Println("error while parsing data:", err)
			}

			icinga := types.NewIcinga("OK - All procs are  OK", warning, critical)

			fmt.Println(icinga)

			fmt.Println("ICINGA DOWN", icinga.Warning.Down, icinga.Warning.DownString)
			fmt.Println("ICINGA UP", icinga.Warning.Up, icinga.Warning.UpString)
			// for _, value := range procs {
			// 	fmt.Println(value)
			// }
		},
	}
)

func init() {
	rootCmd.AddCommand(procsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// procCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	procsCmd.Flags().BoolVar(&count, "count", false, "Count the procs instead of just check the percentage")
}
