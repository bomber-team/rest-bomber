package configuration

import (
	"log"
	"restbomber/discovery/configuration_discovery"
	"restbomber/payloads"
)

/*GlobalConfiguration - config which contains all settings service*/
type GlobalConfiguration struct {
	BomberConfigurationDiscovery *configuration_discovery.ConfigurationDiscovery
	BomberConfigurationService   *payloads.BomberConfig
}

/*SetupBomberConfiguration - setup bomber config service*/
func (glblc *GlobalConfiguration) SetupBomberConfiguration(payload *payloads.BomberConfig) {
	glblc.BomberConfigurationService = payload
}

/*SetupBomberConfigurationDiscovery - setting bomber discovery configuration*/
func (glblc *GlobalConfiguration) SetupBomberConfigurationDiscovery() error {
	glblc.BomberConfigurationDiscovery = &configuration_discovery.ConfigurationDiscovery{}
	if err := glblc.BomberConfigurationDiscovery.ParseFieldFile("config.json"); err != nil {
		log.Println("error while configuration parsed")
		return err
	}
	return nil
}
