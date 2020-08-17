package transitions

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/andygrunwald/go-jira"
	"jirago/lib/client"
)

func Run(key string) (*string, error) {
	transitions, _, jerr := client.Client.Issue.GetTransitions(key)
	if jerr != nil {
		return nil, jerr
	}

	var transitionNames []string

	for _, transition := range transitions {
		key := transition.Name
		transitionNames = append(transitionNames, key)
	}

	var qs = []*survey.Question{
		{
			Name: "transition",
			Prompt: &survey.Select{
				Message: "Choose a transition:",
				Options: transitionNames,
			},
		},
	}

	answers := struct {
		Transition string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return nil, err
	}

	mapNames := make(map[string]jira.Transition)

	for _, transition := range transitions {
		mapNames[transition.Name] = transition
	}

	transition := mapNames[answers.Transition]

	return &transition.ID, nil
}
