package cmd

import (
	"github.com/34N0/rpm-copr/pkg/repos"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List local copr repositories",
	Run: func(cmd *cobra.Command, args []string) {
		repos.ListCoprs(cmd.Flags())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("installed", "i", true, "List installed Copr repositories.")
	listCmd.PersistentFlags().BoolP("enabled", "e", false, "List enabled Copr repositories.")
	listCmd.PersistentFlags().BoolP("disabled", "d", false, "List disabled Copr repositories.")
}
