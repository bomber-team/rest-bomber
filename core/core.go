package core

import (
	"errors"
	"fmt"
	"log"

	"github.com/bomber-team/rest-bomber/models"
	"github.com/bomber-team/rest-bomber/routes/payloads"
)

type (
	/*Core - struct for containing current state of bomber*/
	Core struct {
		Task   *TaskBomber
		Attack *Attacker
		// Generators
		// notifier
		// metrics
	}

	TaskBomber struct {
		CurrentTaskEntry     chan *payloads.Task
		CurrentTask          *payloads.Task
		WorkJob              chan int
		CompletedCurrentTask bool
		CurrentState         *StateBomber
		AttackPreparedInit   chan bool
		Attack               chan bool
	}

	/*StateBomber - state of our bomber*/
	StateBomber struct {
		MemmoryStage          map[string]string
		CurrentStage          int
		CompletedCurrentStage chan bool
		AmountStages          int
		StageWork             chan bool // true - ready to work with one stage // false - working with stage
	}
)

const (
	stageInit = -1
)

const (
	taskIncrementCurrentStage = 1
)

func InitCore() *Core {
	return &Core{
		Task: InitTask()}
}

func InitTask() *TaskBomber {
	return &TaskBomber{
		CurrentTaskEntry:   make(chan *payloads.Task),
		CurrentTask:        nil,
		WorkJob:            make(chan int),
		CurrentState:       InitState(),
		AttackPreparedInit: make(chan bool),
		Attack:             make(chan bool),
	}
}

func InitState() *StateBomber {
	return &StateBomber{
		MemmoryStage:          nil,
		CurrentStage:          stageInit,
		CompletedCurrentStage: make(chan bool),
		AmountStages:          0,
		StageWork:             make(chan bool),
	}
}

func (task *TaskBomber) completeWorkflow() {

}

func (core *Core) preparedAttackMonitor() {
	for {
		<-core.Task.AttackPreparedInit
		query, body := core.Attack.preparingData(core.Task.CurrentTask.Scenario.Stages[core.Task.CurrentState.CurrentStage].StageConfig.AmountRequests)
		fmt.Println(query, body)
	}
}

func (core *Core) attackMonitor() {
	for {
		<-core.Task.Attack
		core.Attack.attack()
	}
}

func (task *TaskBomber) workflowMonitor() {
	for {
		<-task.CurrentState.CompletedCurrentStage
		if err := task.CurrentState.NextStage(); err != nil {
			task.completeWorkflow()
		}
		task.ExecuteStage()
	}
}

func (task *TaskBomber) workApiMonitor() {
	for {
		workJob := <-task.WorkJob
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
		default:
			continue
		}
	}
}

func (task *TaskBomber) taskMonitor() {
	for {
		task.CurrentTask = <-task.CurrentTaskEntry
		task.CurrentState.Clear()
	}
}

func (task *TaskBomber) CurrentTaskAvailable() bool {
	return task.CompletedCurrentTask
}

func (state *StateBomber) Clear() {
	state.CurrentStage = stageInit
}

func (state *StateBomber) NextStage() error {
	if state.CurrentStage+1 < state.AmountStages {
		state.CurrentStage++
		return nil
	} else {
		errors.New("can not increment stage")
	}
}

func (core *Core) AddNewTask(task *payloads.Task) {
	core.Task.CurrentTaskEntry <- task
}

func (task *TaskBomber) ExecuteStage() {
	if !task.CurrentTaskAvailable() {
		currentStage := task.CurrentTask.Scenario.Stages[task.CurrentState.CurrentStage]
		log.Println("start stage: ", currentStage.Name)
		task.AttackPreparedInit <- true
		task.CurrentState.CompletedCurrentStage <- true
	}
}

func (task *TaskBomber) GetCurrentStage() *models.Stage {
	return &task.CurrentTask.Scenario.Stages[task.CurrentState.CurrentStage]
}

func (task *TaskBomber) GetScheme(schemename string) *models.Scheme {
	for _, val := range task.CurrentTask.Schemes {
		if val.ID == schemename {
			return val
		}
	}
	return nil
}

// /*changeCurrentStageIndex - change current stage by last stage + 1*/
// func (bomb *StateBomber) changeCurrentStageIndex() error {
// 	if bomb.CurrentStage+1 < bomb.AmountStages {
// 		bomb.CurrentStage = bomb.CurrentStage + 1
// 		return nil
// 	}
// 	return errors.New("Stage can not be continue, by last stage")
// }

// /*ChangeCurrentStage - change current stage work by index*/
// func (bomb *StateBomber) ChangeCurrentStage() error {
// 	if err := bomb.changeCurrentStageIndex(); err != nil {
// 		log.Println(err.Error())
// 		return err
// 	}
// 	bomb.Stage = &bomb.CurrentScenario.Stages[bomb.CurrentStage]
// 	return nil
// }

// /*GetCurrentStage - get current stage*/
// func (bomb *StateBomber) GetCurrentStage() *models.Stage {
// 	return &bomb.CurrentScenario.Stages[bomb.CurrentStage]
// }

// /*ChangeCurrentState - change current state of bomber*/
// func (bomb *StateBomber) ChangeCurrentState(state *StateBomber) {
// 	bomb.CurrentScenario = state.CurrentScenario
// 	bomb.AmountStages = len(bomb.CurrentScenario.Stages)
// 	bomb.CurrentStage = stageInit
// 	bomb.Stage = nil
// 	bomb.StageWork = make(chan bool, 1)
// 	bomb.ScenarioEnter = make(chan bool, 1)
// }

// /*SchemeOfCurrentStageEmpty - check the current scheme is not empty*/
// func (bomb *StateBomber) SchemeOfCurrentStageEmpty() bool {
// 	if bomb.CurrentScheme == nil {
// 		return true
// 	}
// 	return false
// }

// /*SetupNewEventOfWorkWithOneStage - task emmit start*/
// func (bomb *StateBomber) SetupNewEventOfWorkWithOneStage() {
// 	bomb.StageWork <- true
// }

// /*WaitWorkComplete - */
// func (bomb *StateBomber) WaitWorkComplete() {
// 	<-bomb.StageWork
// }

// /*UpdateMemory - update inline memoy of our bomber*/
// func (core *Core) UpdateMemory() {
// 	core.Memory.Init()
// }

// /*UpdateCore - get the last scenario, update state model*/
// func (core *Core) UpdateCore() {
// 	scenario := core.Memory.GetLastScenario()
// 	core.State.ChangeCurrentState(&StateBomber{
// 		CurrentScenario: scenario,
// 	})
// }

// /*WaitNewScenarioInput - wait scenario enter to core
// return true - if have close core
// return false - if have work with enter scenario
// */
// func (core *Core) WaitNewScenarioInput() bool {
// 	signal := <-core.State.ScenarioEnter
// 	if signal {
// 		return false
// 	}
// 	return true
// }

// /*Convolutional with stages of scenario in current state of our core*/
// func (core *Core) topLayer() {
// 	if sig := core.WaitNewScenarioInput(); sig {
// 		return
// 	}
// 	core.UpdateCore()
// 	core.UpdateMemory()
// 	for i := 0; i < core.State.AmountStages; i++ {
// 		core.SetupNewWork()
// 		if err := core.State.ChangeCurrentStage(); err != nil {
// 			break
// 		}
// 		go func(work chan bool) {

// 		}(core.State.StageWork)
// 	}
// }

// /*InitCore - initialize core bomber*/
// func InitCore() *Core {
// 	core := &Core{}
// 	core.Memory = &MemoryBomber{
// 		Scenaries: nil,
// 		Schemes:   nil,
// 	}

// 	core.State = &StateBomber{
// 		CurrentScheme:   nil,
// 		CurrentStage:    stageInit,
// 		CurrentScenario: nil,
// 	}
// 	return core
// }

// /*SetupNewWork - emmit new event of work*/
// func (core *Core) SetupNewWork() {
// 	core.State.SetupNewEventOfWorkWithOneStage()
// }

// /*WaitWorkStageComplete - wait when work finished*/
// func (core *Core) WaitWorkStageComplete() {
// 	core.State.WaitWorkComplete()
// }

// /*GetLastScenario - get last scenario in queue with scenaries*/
// func (mem *MemoryBomber) GetLastScenario() *models.Scenario {
// 	lastScenarioFromQueue := <-mem.Scenaries
// 	return &lastScenarioFromQueue
// }

// /*Init - initialise buffers*/
// func (mem *MemoryBomber) Init() {
// 	mem.Scenaries = make(chan models.Scenario)
// 	mem.Schemes = make(chan models.Scheme)
// }
