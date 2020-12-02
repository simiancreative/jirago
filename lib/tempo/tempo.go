package tempo

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var client *resty.Client

func init() {
	client = resty.New()
}

func GetWorklogs(id string, start string, stop string) (Worklogs, error) {
	body := worklogReqBody{
		From: start,
		To:   stop,
		ID:   []string{id},
	}

	resp, err := client.R().
		SetBody(body).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+viper.GetString("tempo")).
		Post("https://app.tempo.io/rest/tempo-timesheets/4/worklogs/search")

	if err != nil {
		return nil, err
	}

	worklogs := &Worklogs{}
	json.Unmarshal(resp.Body(), worklogs)

	return *worklogs, nil
}
