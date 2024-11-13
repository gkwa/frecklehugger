package cmd

import (
	"github.com/gkwa/frecklehugger/core"
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:   "notes [path]",
	Short: "Get git notes from repository",
	Long:  `Retrieves all git notes from the specified repository path`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		logger := LoggerFrom(ctx)
		logger.Info("Running notes command")
		core.PrintNotes(ctx, args[0], logger)
	},
}

func init() {
	rootCmd.AddCommand(notesCmd)
}
