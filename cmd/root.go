/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	filter   string
	host     string
	port     int
	secure   bool
	token    string
	perfData bool

	warning  string
	critical string

	rootCmd = &cobra.Command{
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
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Defines the filter for the request. A dot (.) Means no filter at all")
	rootCmd.PersistentFlags().IntVar(&port, "port", 20480, "Defines the filter for the request. A dot (.) Means no filter at all")
	rootCmd.PersistentFlags().StringVar(&token, "token", ".", "Defines the filter for the request. A dot (.) Means no filter at all")
	rootCmd.PersistentFlags().BoolVar(&secure, "secure", false, "Defines the filter for the request. A dot (.) Means no filter at all")

	rootCmd.PersistentFlags().BoolVar(&perfData, "perf", false, "Defines if perfData is added to the command")

	rootCmd.PersistentFlags().StringVar(&filter, "filter", "", "Defines the filter for the request. A dot (.) Means no filter at all")

	if filter == "." {
		filter = ""
	}
}
