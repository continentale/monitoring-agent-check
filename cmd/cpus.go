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
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.AddCommand(cpusCmd)

	cpusCmd.PersistentFlags().BoolVar(&perCPU, "perCPU", false, "")
}
