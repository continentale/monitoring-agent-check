/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// memCmd represents the mem command
var (
	convert string
	memCmd  = &cobra.Command{
		Use:   "mem",
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
	rootCmd.AddCommand(memCmd)

	memCmd.Flags().StringVar(&convert, "convert", "GB", "The unit in which the data is converted. Supported is KB, MB, GB, TB")
}
