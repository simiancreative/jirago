package loe

import (
	"github.com/AlecAivazis/survey/v2"
)

func Run() (*string, error) {
	var qs = []*survey.Question{
		{
			Name:      "loe",
			Prompt:    &survey.Input{Message: "Enter an LoE"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}

	answers := struct {
		Loe string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		return nil, err
	}

	return &answers.Loe, nil
}
