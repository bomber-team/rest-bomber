package payloads

import "github.com/bomber-team/rest-bomber/models"

/*ScenarioPayload - payload for addding scenario for attac*/
type ScenarioPayload struct {
	Scenario *models.Scenario `json:"scenario"`
}
