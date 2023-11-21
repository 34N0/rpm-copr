package cmd

import (
	"github.com/34N0/rpm-copr/pkg/repos"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List local copr repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repos.ListCoprs()
	},
}

func init() { rootCmd.AddCommand(listCmd) }
