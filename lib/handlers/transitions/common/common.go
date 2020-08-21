package common

import (
	"jirago/lib/client"
	"jirago/lib/handlers"
)

var Handler = handlers.Handler{
	Run: func(params handlers.Params) error {
		issueId := params["issue"].(string)
		transitionId := params["transition"].(string)

		_, err := client.Client.Issue.DoTransition(issueId, transitionId)

		if err != nil {
			return err
		}

		return nil
	},
}
