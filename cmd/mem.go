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
		Short: "checks mem values of the target",
		Long:  `Checks the return values of the agent from the mem endpoint`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	RootCmd.AddCommand(memCmd)
}
