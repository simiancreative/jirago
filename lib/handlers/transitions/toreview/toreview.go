package toreview

import (
	"jirago/lib/client"
	"jirago/lib/handlers"
	"jirago/lib/surveys/issues/timelogged"
)

var Handler = handlers.Handler{
	Run: func(params handlers.Params) error {
		timelogged, lerr := timelogged.Run()
		if lerr != nil {
			return lerr
		}

		issueId := params["issue"].(string)
		transitionId := params["transition"].(string)

		var payload = map[string]interface{}{
			"transition": map[string]string{
				"id": transitionId,
			},
			"fields": map[string]interface{}{
				"resolution": map[string]string{
					"name": timelogged.Resolution,
				},
			},
			"update": map[string]interface{}{
				"worklog": []map[string]interface{}{
					{"add": map[string]string{
						"timeSpent": timelogged.Time,
					}},
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
