/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var (
	perCPU  bool
	cpusCmd = &cobra.Command{
		Use:   "cpus",
		Short: "checks cpu values of the target",
		Long: `checks the return values of the agent from the cpu endpoint

		example output from the agent with perCPU = false: 
		[
			{
				"cpu": "cpu-total",
				"user": 94.32,
				"system": 76.49,
				"idle": 17205.19,
				"nice": 0.17,
				"iowait": 18.29,
				"irq": 0,
				"softirq": 33.28,
				"steal": 0,
				"guest": 0,
				"guestNice": 0
			}
		]
		
		Now you can check if you have enough cpu resources available with the cpuusage command or check the load on unix systems with the load command

		./monitoring-agent cpus usage --help
		./monitoring-agent cpus load  --help`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	RootCmd.AddCommand(cpusCmd)

	cpusCmd.PersistentFlags().BoolVar(&perCPU, "perCPU", false, "")
}
