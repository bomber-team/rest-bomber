package payloads

import "restbomber/models"

/*ScenarioPayload - payload for addding scenario for attac*/
type ScenarioPayload struct {
	Scenario *models.Scenario `json:"scenario"`
}
