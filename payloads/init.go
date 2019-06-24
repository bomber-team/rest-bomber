package payloads

/*InitPayload - */
type InitPayload struct {
	Field string `json:"init-work"`
}

const (
	// ActionStartWork - starting routine for requesting
	ActionStartWork = "start-work"
	// ActionStopWork - stoping routing for requesting
	ActionStopWork = "stop-work"
)

const (
	ActionStartWorkID          = 0
	ActionStopWorkID           = 1
	ActionWriteConfigurationID = 2
	ActionWriteSchemeID        = 3
	ActionWriteScenarioID      = 4
	ActionAttackerStartID      = 5
)
