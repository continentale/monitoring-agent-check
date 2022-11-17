/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var (
	filter   string
	host     string
	port     int
	secure   bool
	token    string
	perfData bool
	verbose  bool
	warning  string
	critical string
	onValue  string
	mode     string

	RootCmd = &cobra.Command{
		Use:   "monitoring-agent-check",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Defines the filter for the request. A dot (.) Means no filter at all")
	RootCmd.PersistentFlags().IntVar(&port, "port", 20480, "Defines the filter for the request. A dot (.) Means no filter at all")
	RootCmd.PersistentFlags().StringVar(&token, "token", ".", "Defines the filter for the request. A dot (.) Means no filter at all")
	RootCmd.PersistentFlags().BoolVar(&secure, "secure", false, "Defines the filter for the request. A dot (.) Means no filter at all")
	RootCmd.PersistentFlags().StringVar(&warning, "warning", "90", "The value on which field the value is checked")
	RootCmd.PersistentFlags().StringVar(&critical, "critical", "95", "The value on which field the value is checked")
	RootCmd.PersistentFlags().BoolVar(&perfData, "perf", true, "Defines if perfData is added to the command")
	RootCmd.PersistentFlags().StringVar(&filter, "filter", "", "Defines the filter for the request. A dot (.) Means no filter at all")
	RootCmd.PersistentFlags().StringVar(&mode, "mode", "", "Defines the filter for the request. A dot (.) Means no filter at all")

	RootCmd.PersistentFlags().StringVar(&onValue, "on", "available", "The value on which field the value is checked")

	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Prints out debug messages for developing")

	if filter == "." {
		filter = ""
	}
}
