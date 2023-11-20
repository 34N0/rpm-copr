package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable the name/project Copr repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disable called")
	},
}

func init() { rootCmd.AddCommand(disableCmd) }
