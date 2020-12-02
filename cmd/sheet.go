package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/spf13/cobra"

	"jirago/lib/client"
	"jirago/lib/tempo"
)

// getSheetCommand represents the transition command
var getSheetCommand = &cobra.Command{
	Use:   "sheet [natural language start time] [natural language end time]",
	Short: "get a tempo time sheet for period",
	Long:  `get a tempo time sheet for period`,
	Args:  cobra.MinimumNArgs(2),
	RunE:  getSheetRun,
}

func init() {
	rootCmd.AddCommand(getSheetCommand)
}

func getSheetRun(cmd *cobra.Command, args []string) error {
	var w = when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	start, err := w.Parse(args[0], time.Now())
	if err != nil {
		return err
	}
	if start == nil {
		return fmt.Errorf("No Matches for '%s'", args[0])
	}

	end, endErr := w.Parse(args[1], time.Now())
	if endErr != nil {
		return endErr
	}

	if end == nil {
		return fmt.Errorf("No Matches for '%s'", args[1])
	}

	startStr := fmt.Sprintf("%d-%02d-%02d",
		start.Time.Year(), start.Time.Month(), start.Time.Day(),
	)
	endStr := fmt.Sprintf("%d-%02d-%02d",
		end.Time.Year(), end.Time.Month(), end.Time.Day(),
	)

	user, _, reqErr := client.Client.User.GetSelf()
	if reqErr != nil {
		return reqErr
	}

	worklogs, _ := tempo.GetWorklogs(user.AccountID, startStr, endStr)

	rows := []table.Row{}
	for i, wl := range worklogs {
		rows = append(rows, []interface{}{
			i, wl.Started, wl.TimeSpent, wl.Issue.Key,
		})
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRows(rows)
	t.Render()

	return nil
}
