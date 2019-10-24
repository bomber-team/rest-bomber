package payloads

import "restbomber/models"

/*SchemePayload - payload for configuring datas models*/
type SchemePayload struct {
	Schemes []*models.Scheme `json:"schemes"`
}
