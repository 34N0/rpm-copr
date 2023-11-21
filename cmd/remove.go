package cmd

import (
	"github.com/34N0/rpm-copr/pkg/repos"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the copr repository",
	Run: func(cmd *cobra.Command, args []string) {
		repos.NewCopr(args).Remove()
	},
}

func init() { rootCmd.AddCommand(removeCmd) }
