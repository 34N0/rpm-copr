package cmd

import (
	"github.com/34N0/rpm-copr/pkg/repos"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable the copr repository",
	Run: func(cmd *cobra.Command, args []string) {
		repos.NewCopr(args).Disable()
	},
}

func init() { rootCmd.AddCommand(disableCmd) }
