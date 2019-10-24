package payloads

/*BomberConfig - main configuration for bomber*/
type BomberConfig struct {
	MetricsAddress         string `json:"metricsAddress"` // address which will be using in metric sender
	ServerAddress          string `json:"server_address"` // address for main backend bomber-server
	NotifyAllActionChanges bool   `json:"notify_all_action"`
}
