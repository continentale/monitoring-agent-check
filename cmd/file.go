/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "checks a filecontent or file age from the target",
	Long: `checks the content or the file stats from a given file. 
	Example output for contentOnly (to check the content)
	
	PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin"
	PATH=$PATH:/usr/local/go/bin
	
	example output for a file
	{
		"Path": "/etc/fstab",
		"IsDir": false,
		"ModTime": 1629409395,
		"Mode": "-rw-r--r--",
		"Name": "fstab",
		"Size": 43,
		"Content": "LABEL=cloudimg-rootfs\t/\t ext4\tdefaults\t0 1\n"
	}

	./monitoring-agent file content --help
	./monitoring-agent file stats --help
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

func init() {
	RootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
