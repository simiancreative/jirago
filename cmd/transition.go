package cmd

import (
	"github.com/spf13/cobra"

	actions "jirago/lib/handlers/transitions"
	// "jirago/lib/logger"
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

	// logger.Info("---", logger.Fields{"t": t})

	action := actions.Actions[t]

	var params = map[string]interface{}{
		"issue":      i,
		"transition": t,
	}
	aerr := action.Run(params)
	if aerr != nil {
		return aerr
	}

	return nil
}
