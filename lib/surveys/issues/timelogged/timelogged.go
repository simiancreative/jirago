package timelogged

import (
	"github.com/AlecAivazis/survey/v2"
)

type Answer struct {
	Resolution string
	Time       string
}

func Run() (*Answer, error) {
	var qs = []*survey.Question{
		{
			Name:      "time",
			Prompt:    &survey.Input{Message: "Enter time spent"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "resolution",
			Prompt: &survey.Select{
				Message: "Choose a resolution",
				Options: []string{
					"Fixed",
					"Won't Fix",
					"Duplicate",
					"Incomplete",
					"Cannot Reproduce",
					"Done",
					"Won't Do",
				},
			},
		},
	}

	var answers = Answer{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return nil, err
	}

	return &answers, nil
}
