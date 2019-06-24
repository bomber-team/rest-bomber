package models

//Configuration for attacker
type Configuration struct {
	AttackAddress      string `json:"attack-address"`
	BufferSize         int    `json:"buffer-size"`
	Type               string `json:"type"`
	TimeBetweenAttacks string `json:"time-between-attacks"`
	AmountAttacks      int    `json:"amount-attacks"`
	ScenarioID         string `json:"scenario-id"`
	LogLevel           string `json:"log_level"`
	ApplicationType    string `json:"application-type"`
}

/*CheckValid - */
func (conf *Configuration) CheckValid() bool {
	return conf.ValidationAttackAddress() &&
		conf.ValidationScenarioID()
}

/*ValidationAttackAddress - */
func (conf *Configuration) ValidationAttackAddress() bool {
	if len(conf.AttackAddress) > 0 {
		return true
	} else {
		return false
	}
}

/*ValidationScenarioID - */
func (conf *Configuration) ValidationScenarioID() bool {
	if len(conf.ScenarioID) > 0 {
		return true
	} else {
		return false
	}
}
