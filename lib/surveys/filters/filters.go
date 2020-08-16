package filters

import (
	"github.com/AlecAivazis/survey/v2"
	"jirago/lib/client"
	"jirago/lib/logger"
)

func Run() *string {
	filters, _, _ := client.Client.Filter.GetMyFilters(nil)
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
		logger.Error("---", logger.Fields{"err": err.Error()})
		return nil
	}

	mapNames := make(map[string]client.Filter)

	for _, v := range filters {
		mapNames[v.Name] = v
	}

	filter := mapNames[answers.Filter]

	return &filter.ID
}
