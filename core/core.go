package core

import (
	"errors"
	"log"

	"github.com/bomber-team/rest-bomber/models"
	"github.com/bomber-team/rest-bomber/routes/payloads"
)

type (
	/*MemoryBomber - memory which contain all schemes, scenarius*/
	MemoryBomber struct {
		Scenaries chan models.Scenario
		Schemes   chan models.Scheme
		WorkJob   chan int
	}

	/*Core - struct for containing current state of bomber*/
	Core struct {
		Memory *MemoryBomber
		State  *StateBomber
	}

	/*StateBomber - state of our bomber*/
	StateBomber struct {
		CurrentScheme   *models.Scheme
		CurrentScenario *models.Scenario
		CurrentStage    int
		AmountStages    int
		Stage           *models.Stage
		StageWork       chan bool // true - ready to work with one stage // false - working with stage
		ScenarioEnter   chan bool // true - insert new scenario // false - remove core
	}
)

const (
	scenarioEmpty = -1
	stageInit     = -1
)

func (core *Core) handlingRoutineActions() {
	log.Print("worker for device starting")
	for {
		var workJob int = <-core.Memory.WorkJob
		log.Print("in channel handle new value")
		switch workJob {
		case payloads.ActionStartWorkID:
			log.Print("start work")
			break
		case payloads.ActionStopWorkID:
			log.Print("stop work")
			break
		case payloads.ActionRestartWorkID:
			log.Print("Restart work")
			break
		// case payloads.ActionWriteConfigurationID:
		// 	log.Print("start configration build")
		// 	break
		// case payloads.ActionWriteScenarioID:
		// 	log.Print("start scenario build")
		// 	break
		// case payloads.ActionWriteSchemeID:
		// 	log.Print("start scheme build")
		// 	break
		default:
			continue
		}
	}
}

/*changeCurrentStageIndex - change current stage by last stage + 1*/
func (bomb *StateBomber) changeCurrentStageIndex() error {
	if bomb.CurrentStage+1 < bomb.AmountStages {
		bomb.CurrentStage = bomb.CurrentStage + 1
		return nil
	}
	return errors.New("Stage can not be continue, by last stage")
}

/*ChangeCurrentStage - change current stage work by index*/
func (bomb *StateBomber) ChangeCurrentStage() error {
	if err := bomb.changeCurrentStageIndex(); err != nil {
		log.Println(err.Error())
		return err
	}
	bomb.Stage = &bomb.CurrentScenario.Stages[bomb.CurrentStage]
	return nil
}

/*GetCurrentStage - get current stage*/
func (bomb *StateBomber) GetCurrentStage() *models.Stage {
	return &bomb.CurrentScenario.Stages[bomb.CurrentStage]
}

/*ChangeCurrentState - change current state of bomber*/
func (bomb *StateBomber) ChangeCurrentState(state *StateBomber) {
	bomb.CurrentScenario = state.CurrentScenario
	bomb.AmountStages = len(bomb.CurrentScenario.Stages)
	bomb.CurrentStage = stageInit
	bomb.Stage = nil
	bomb.StageWork = make(chan bool, 1)
	bomb.ScenarioEnter = make(chan bool, 1)
}

/*SchemeOfCurrentStageEmpty - check the current scheme is not empty*/
func (bomb *StateBomber) SchemeOfCurrentStageEmpty() bool {
	if bomb.CurrentScheme == nil {
		return true
	}
	return false
}

/*SetupNewEventOfWorkWithOneStage - task emmit start*/
func (bomb *StateBomber) SetupNewEventOfWorkWithOneStage() {
	bomb.StageWork <- true
}

/*WaitWorkComplete - */
func (bomb *StateBomber) WaitWorkComplete() {
	<-bomb.StageWork
}

/*UpdateMemory - update inline memoy of our bomber*/
func (core *Core) UpdateMemory() {
	core.Memory.Init()
}

/*UpdateCore - get the last scenario, update state model*/
func (core *Core) UpdateCore() {
	scenario := core.Memory.GetLastScenario()
	core.State.ChangeCurrentState(&StateBomber{
		CurrentScenario: scenario,
	})
}

/*WaitNewScenarioInput - wait scenario enter to core
return true - if have close core
return false - if have work with enter scenario
*/
func (core *Core) WaitNewScenarioInput() bool {
	signal := <-core.State.ScenarioEnter
	if signal {
		return false
	}
	return true
}

/*Convolutional with stages of scenario in current state of our core*/
func (core *Core) topLayer() {
	if sig := core.WaitNewScenarioInput(); sig {
		return
	}
	core.UpdateCore()
	core.UpdateMemory()
	for i := 0; i < core.State.AmountStages; i++ {
		core.SetupNewWork()
		if err := core.State.ChangeCurrentStage(); err != nil {
			break
		}
		go func(work chan bool) {

		}(core.State.StageWork)
	}
}

/*InitCore - initialize core bomber*/
func (core *Core) InitCore() *Core {
	core.Memory = &MemoryBomber{
		Scenaries: nil,
		Schemes:   nil,
	}

	core.State = &StateBomber{
		CurrentScheme:   nil,
		CurrentStage:    stageInit,
		CurrentScenario: nil,
	}
	return core
}

/*SetupNewWork - emmit new event of work*/
func (core *Core) SetupNewWork() {
	core.State.SetupNewEventOfWorkWithOneStage()
}

/*WaitWorkStageComplete - wait when work finished*/
func (core *Core) WaitWorkStageComplete() {
	core.State.WaitWorkComplete()
}

/*GetLastScenario - get last scenario in queue with scenaries*/
func (mem *MemoryBomber) GetLastScenario() *models.Scenario {
	lastScenarioFromQueue := <-mem.Scenaries
	return &lastScenarioFromQueue
}

/*Init - initialise buffers*/
func (mem *MemoryBomber) Init() {
	mem.Scenaries = make(chan models.Scenario)
	mem.Schemes = make(chan models.Scheme)
}
