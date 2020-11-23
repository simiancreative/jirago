package cmd

import (
	"fmt"
	"net/http"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"

	"jirago/lib/client"
)

// addTimeCommand represents the transition command
var addTimeCommand = &cobra.Command{
	Use:   "time",
	Short: "Add time to jira tickets",
	RunE:  addTimeRun,
}

func init() {
	rootCmd.AddCommand(addTimeCommand)
}

func addTimeRun(cmd *cobra.Command, args []string) error {

	issueID := args[0]
	loe := args[1]
	cb := func(*http.Request) error {
		return nil
	}

	worklog := jira.WorklogRecord{
		IssueID:   issueID,
		TimeSpent: loe,
	}

	_, _, err := client.Client.Issue.AddWorklogRecord(issueID, &worklog, cb)

	if err != nil {
		return err
	}

	fmt.Printf("\nAdded %s to %s\n", loe, issueID)

	return nil
}
