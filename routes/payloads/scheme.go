package payloads

import "github.com/bomber-team/rest-bomber/models"

/*SchemePayload - payload for configuring datas models*/
type SchemePayload struct {
	Schemes []*models.Scheme `json:"schemes"`
}
