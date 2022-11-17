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
		Short: "checks disks values of the target",
		Long: `checks the return values of the agent from the cpu endpoint

		example output from the agent
		
		[
			{
				"usage": {
					"path": "/",
					"fstype": "ext2/ext3",
					"total": 269490393088,
					"free": 231327899648,
					"used": 24401821696,
					"usedPercent": 9.54203585244415,
					"inodesTotal": 16777216,
					"inodesUsed": 880465,
					"inodesFree": 15896751,
					"inodesUsedPercent": 5.247980356216431
				},
				"details": {
					"device": "/dev/sdb",
					"mountpoint": "/",
					"fstype": "ext4",
					"opts": [
						"rw",
						"relatime"
					]
				}
			}
		]

		Now you can check if you have enough disk space, temp space or enough inodes

		./monitoring-agent disks inodes --help
		./monitoring-agent disks temp --help
		./monitoring-agent disks usage --help`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	RootCmd.AddCommand(disksCmd)

	// disksCmd.Flags().StringVar(&onValue, "on", "usage.usedPercent", "The value on which field the value is checked")
}
