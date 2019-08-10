package core

import "rest-bomber/models"

type (
	/*MemoryBomber - memory which contain all schemes, scenarius*/
	MemoryBomber struct {
		Scenaries []chan models.Scenario
		Schemes   []chan models.Scheme
	}

	/*Core - struct for containing current state of bomber*/
	Core struct {
		Memory *MemoryBomber
		State  *StateBomber
	}

	/*StateBomber - state of our bomber*/
	StateBomber struct {
		CurrentSchemes  []int
		CurrentScenario int
		CurrentStage    int
	}
)

const (
	scenarioEmpty = -1
	stageInit     = 0
	stageReady    = 1
	stageWork     = 2
)

/*InitCore - initialize core bomber*/
func (core *Core) InitCore() *Core {
	core.Memory = &MemoryBomber{
		Scenaries: nil,
		Schemes:   nil,
	}

	core.State = &StateBomber{
		CurrentSchemes:  nil,
		CurrentStage:    stageInit,
		CurrentScenario: scenarioEmpty,
	}
	return core
}
