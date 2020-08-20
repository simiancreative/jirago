package beginwork

import (
	"github.com/andygrunwald/go-jira"

	"jirago/lib/client"
	"jirago/lib/handlers"
	"jirago/lib/surveys/issues/loe"
)

var Handler = handlers.Handler{
	Run: func(params handlers.Params) error {
		estimate, lerr := loe.Run()
		if lerr != nil {
			return lerr
		}

		issueId := params["issue"]
		transitionId := params["transition"]
		var payload = map[string]interface{}{
			"transition": transitionId,
			"fields": &jira.IssueFields{
				TimeTracking: &jira.TimeTracking{
					OriginalEstimate: *estimate,
				},
			},
		}
		_, err := client.Client.Issue.DoTransitionWithPayload(issueId, payload)

		if err != nil {
			return err
		}

		return nil
	},
}
