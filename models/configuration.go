package models

//Configuration for attacker
type Configuration struct {
	AttackAddress      string `json:"attack-address"`
	ServerAddress      string `json:"server-address"`
	BufferSize         int    `json:"buffer-size"`
	Type               string `json:"type"`
	TimeBetweenAttacks string `json:"time-between-attacks"`
	AmountAttacks      int    `json:"amount-attacks"`
	PathScenarioFile   string `json:"path-scenario-file"`
	LogLevel           string `json:"log_level"`
	ApplicationType    string `json:"application-type"`
}
