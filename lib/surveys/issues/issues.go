package issues

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/andygrunwald/go-jira"
	"jirago/lib/client"
	"jirago/lib/logger"
)

func makeKey(issue jira.Issue) string {
	return fmt.Sprintf("%s - %s ", issue.Key, issue.Fields.Summary)
}

func Run(jql string) (*string, error) {
	issues, _, err := client.Client.Issue.Search(jql, &jira.SearchOptions{
		Fields: []string{"summary"},
	})
	if err != nil {
		return nil, err
	}

	var issueNames []string

	for _, issue := range issues {
		key := makeKey(issue)
		issueNames = append(issueNames, key)
	}

	var qs = []*survey.Question{
		{
			Name: "issue",
			Prompt: &survey.Select{
				Message: "Choose an issue:",
				Options: issueNames,
			},
		},
	}

	answers := struct {
		Issue string
	}{}

	if err := survey.Ask(qs, &answers); err != nil {
		logger.Error("---", logger.Fields{"err": err.Error()})
		return nil, err
	}

	mapNames := make(map[string]jira.Issue)

	for _, issue := range issues {
		key := makeKey(issue)
		mapNames[key] = issue
	}

	issue := mapNames[answers.Issue]

	return &issue.Key, nil
}
