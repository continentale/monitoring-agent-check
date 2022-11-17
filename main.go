/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/continentale/monitoring-agent-check/cmd"
)

func main() {
	/*
		err := doc.GenMarkdownTree(cmd.RootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	*/

	cmd.Execute()
}
