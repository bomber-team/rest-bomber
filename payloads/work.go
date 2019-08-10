package payloads

/*WorkPayload - payload for manipulate with work bomber*/
type WorkPayload struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}

const (
	/*ActionStartPreparingAttack - action for starting preparing attack by scheme
	 */
	ActionStartPreparingAttack = 0
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
