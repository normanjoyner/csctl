package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/containership/csctl/pkg/buildinfo"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the current version of csctl",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(buildinfo.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
