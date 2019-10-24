package payloads

import "github.com/bomber-team/rest-bomber/models"

/*Task - setup task for bomber work*/
type Task struct {
	Scenario *models.Scenario `json:"scenario"`
	Schemes  []*models.Scheme `json:"schemes"`
}
