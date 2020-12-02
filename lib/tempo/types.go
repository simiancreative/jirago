package tempo

import (
	"github.com/simplereach/timeutils"
)

type worklogReqBody struct {
	From string   `json:"from"`
	To   string   `json:"to"`
	ID   []string `json:"workerId"`
}

type Issue struct {
	Summary string `json:"summary"`
	Key     string `json:"key"`
	Status  string `json:"issueStatus"`
}

type Worklog struct {
	TimeSpent string         `json:"timeSpent"`
	Started   timeutils.Time `json:"started"`
	Issue     Issue          `json:"issue"`
}

type Worklogs []Worklog
