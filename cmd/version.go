package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gkwa/frecklehugger/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of frecklehugger",
	Long:  `All software has versions. This is frecklehugger's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
