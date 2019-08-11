package requesters

import "rest-bomber/models"

type ConfigurationRequest struct {
	ServerAddress string
}

func (requester *ConfigurationRequest) StartRequest() *models.Configuration {
	return nil
}
