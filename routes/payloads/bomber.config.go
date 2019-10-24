package payloads

/*BomberConfig - main configuration for bomber*/
type BomberConfig struct {
	ID             string `json:"id"`             // identificator of our bomber
	MetricsAddress string `json:"metricsAddress"` // address which will be using in metric sender
	ServerAddress  string `json:"server_address"` // address for main backend bomber-server
}
