package cmd

import (
	"github.com/34N0/rpm-copr/pkg/repos"
	"github.com/spf13/cobra"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable the copr repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repos.NewCopr(args).Enable()
	},
}

func init() { rootCmd.AddCommand(enableCmd) }
