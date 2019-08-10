package core

import "rest-bomber/models"

/*Attack - module which convolutional commands for atack and another*/
type Attack struct {
	ActionForAttacker chan int // action for work stage (stop, start, restart)
	Schemas           []models.Scheme
	Scenario          models.Scenario
}

func (atck *Attack) worker() {
	for {
		action := <-atck.ActionForAttacker
		atck.convolutional(action)
	}
}

func (atck *Attack) convolutional(action int) {
	switch action {
	case 0:
	}
}

func (atck *Attack) startAttack() {

}

func (atck *Attack) stopAttack() {

}

func (atck *Attack) restartAttack() {

}

func (atck *Attack) setupScenario(scenario *models.Scenario) {
	atck.Scenario = *scenario
}

func (atck *Attack) setupSchemas(schemsa []models.Scheme) {
	atck.Schemas = schemsa
}

// setup schemas
// setup scenario
