package models

type (
	/*Scenario - main model for scenario stages*/
	Scenario struct {
		Stages        []Stage               `json:"stages"`
		Configuration ConfigurationScenario `json:"configuration"`
	}

	/*Stage - model for stage*/
	Stage struct {
		Name        string             `json:"name"`
		Scheme      string             `json:"scheme"`
		Address     string             `json:"address"`
		StageConfig StageConfiguration `json:"configuration"`
	}

	/*StageConfiguration - config for stage*/
	StageConfiguration struct {
		AmountRequests      int  `json:"amountRequests"`
		TimeoutOnOneRequest int `json:"timeout_one_request"`
		TimeBetweenAttacks  int  `json:"timebetweenattacs"`
		NotifyAfterComplete bool `json:"notifiAfterComplete"`
		SendMetrics         bool `json:"sendMetrics"`
		Log                 bool `json:"log"`
		GeneratedStash      bool `json:"generatedStash"`
	}

	/*ConfigurationScenario - configure for scenario*/
	ConfigurationScenario struct {
		Replay             int  `json:"replay"`
		ReplayAfterFailed  bool `json:"replayAfterFailed"`
		ReplayAfterTimeout int  `json:"replayAfterTimeout"`
	}
)
