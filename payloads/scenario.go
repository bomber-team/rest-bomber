package payloads

import "rest-bomber/models"

/*ScenarioPayload - payload for addding scenario for attac*/
type ScenarioPayload struct {
	Scenario *models.Scenario `json:"scenario"`
}
