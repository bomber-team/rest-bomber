package requesters

import "rest-bomber/models"

type SchemeRequester struct {
	ServerAddress string
}

func (requester *SchemeRequester) StartRequest() *models.Scheme {
	return nil
}
