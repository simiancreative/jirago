package filters

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/andygrunwald/go-jira"
	"jirago/lib/client"
)

func Run() (*string, error) {
	filters, _, ferr := client.Client.Filter.GetMyFilters(nil)
	if ferr != nil {
		return nil, ferr
	}
	var filterNames []string

	for _, filter := range filters {
		filterNames = append(filterNames, filter.Name)
	}

	var qs = []*survey.Question{
		{
			Name: "filter",
			Prompt: &survey.Select{
				Message: "Choose a filter:",
				Options: filterNames,
			},
		},
	}

	answers := struct {
		Filter string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return nil, err
	}

	mapNames := make(map[string]*jira.Filter)

	for _, v := range filters {
		mapNames[v.Name] = v
	}

	filter := mapNames[answers.Filter]

	return &filter.Jql, nil
}
