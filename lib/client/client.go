package client

import (
	"github.com/andygrunwald/go-jira"
	"github.com/spf13/viper"
)

var (
	Client *jira.Client
)

func Setup() {
	tp := jira.BasicAuthTransport{
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}

	client, _ := jira.NewClient(tp.Client(), viper.GetString("host"))
	Client = client
}
