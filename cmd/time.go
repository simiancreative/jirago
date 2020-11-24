package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/spf13/cobra"

	"jirago/lib/client"
)

// addTimeCommand represents the transition command
var addTimeCommand = &cobra.Command{
	Use:   "time [issueID] [duration] <natural language start time>",
	Short: "Add worklogs to jira tasks",
	Long: `Add worklogs to jira tasks.

Arguments

Issue Id            - https://bit.ly/33dYq5g
Jira Duration       - https://bit.ly/3nXvl5X
Optional Start Time - https://bit.ly/2KDYVyX
	`,
	Args: cobra.MinimumNArgs(2),
	RunE: addTimeRun,
}

func init() {
	rootCmd.AddCommand(addTimeCommand)
}

func addTimeRun(cmd *cobra.Command, args []string) error {
	var w = when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	issueID := args[0]
	loe := args[1]
	cb := func(*http.Request) error {
		return nil
	}

	worklog := jira.WorklogRecord{
		IssueID:   issueID,
		TimeSpent: loe,
	}

	if len(args) == 3 {
		resp, err := w.Parse(args[2], time.Now())
		if err != nil {
			return err
		}

		if resp == nil {
			return fmt.Errorf("No Matches for '%s'", args[2])
		}

		started := jira.Time(resp.Time)
		worklog.Started = &started
	}

	_, _, err := client.Client.Issue.AddWorklogRecord(issueID, &worklog, cb)

	if err != nil {
		return err
	}

	if worklog.Started != nil {
		fmt.Printf("\nAdded %s to %s at %s\n", loe, issueID, time.Time(*worklog.Started))
		return nil
	}

	fmt.Printf("\nAdded %s to %s\n", loe, issueID)

	return nil
}
