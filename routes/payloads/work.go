package payloads

/*WorkPayload - payload for manipulate with work bomber*/
type WorkPayload struct {
	Action int `json:"action"`
}

const (
	/*ActionStartWorkID - action for starting work by id
	 */
	ActionStartWorkID = 1
	/*ActionStopWorkID - action for starting work by id
	 */
	ActionStopWorkID = 2
	/*ActionRestartWorkID - action for starting work by id
	 */
	ActionRestartWorkID = 3
)
