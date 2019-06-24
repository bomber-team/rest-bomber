package payloads

/*AttackPayload - payload for action with attacker*/
type AttackPayload struct {
	Field string `json:"attack-action"`
}

const (
	// StartAttack - request payload for activating bomber
	StartAttack = "start-attack"
)
