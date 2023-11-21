/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "rpm-copr",
	Version: "v0.8-alpha",
	Short:   "Interact with Copr repositories.",
	Long:    `rpm-copr is a Command Line Interface that ports the COPR dnf command to immutable (OSTree) images.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
