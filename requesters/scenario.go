package requesters

import (
	"rest-bomber/models"
)

type ScenarioRequester struct {
	ServerAddress string
}

func (requester *ScenarioRequester) StartRequest() *models.Scenario {
	return nil
}
