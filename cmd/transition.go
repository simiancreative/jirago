package cmd

import (
	"github.com/spf13/cobra"

	"jirago/lib/logger"
	"jirago/lib/surveys/filters"
	"jirago/lib/surveys/issues"
	"jirago/lib/surveys/issues/transitions"
)

// transitionCmd represents the transition command
var transitionCmd = &cobra.Command{
	Use:   "transition",
	Short: "Transition jira tickets",
	RunE:  run,
}

func init() {
	rootCmd.AddCommand(transitionCmd)
}

func run(cmd *cobra.Command, args []string) error {
	// Get filter jql
	filterJql, ferr := filters.Run()
	if ferr != nil {
		return ferr
	}
	f := *filterJql

	// Get issue key
	issueKey, ierr := issues.Run(f)
	if ierr != nil {
		return ierr
	}
	i := *issueKey

	// Get issue transition
	tid, err := transitions.Run(i)
	if err != nil {
		return err
	}
	t := *tid

	logger.Info("---", logger.Fields{"transition": t})

	return nil
}
