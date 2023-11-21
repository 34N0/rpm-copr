/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/34N0/rpm-copr/pkg/copr"
	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable the name/project Copr repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		copr.NewCopr(args).Enable()
	},
}

func init() { rootCmd.AddCommand(enableCmd) }
