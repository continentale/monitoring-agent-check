/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var (
	disksCmd = &cobra.Command{
		Use:   "disks",
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
	rootCmd.AddCommand(disksCmd)

	// disksCmd.Flags().StringVar(&onValue, "on", "usage.usedPercent", "The value on which field the value is checked")
}
