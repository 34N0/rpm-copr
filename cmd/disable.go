package cmd

import (
	"github.com/34N0/rpm-copr/pkg/repo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable the name/project Copr repository",
	Run: func(cmd *cobra.Command, args []string) {
		repo.NewCopr(args).Disable()
	},
}

func init() { rootCmd.AddCommand(disableCmd) }
