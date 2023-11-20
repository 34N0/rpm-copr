/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable the name/project Copr repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enable called")
	},
}

func init() { rootCmd.AddCommand(enableCmd) }
