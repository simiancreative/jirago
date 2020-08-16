package cmd

import (
	"github.com/spf13/cobra"
	"jirago/lib/logger"
	"jirago/lib/surveys/filters"
)

// transitionCmd represents the transition command
var transitionCmd = &cobra.Command{
	Use:   "transition",
	Short: "Transition jira tickets",
	Run: func(cmd *cobra.Command, args []string) {
		filterID := filters.Run()
		i := *filterID
		logger.Info("---", logger.Fields{"filter": i})

	},
}

func init() {
	rootCmd.AddCommand(transitionCmd)
}
