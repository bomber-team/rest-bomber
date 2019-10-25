package core

/*Attacker - main entry for attack start*/
type Attacker struct {
	Task           *TaskBomber      // contain current schemas and scenario
	Generators 	generators.Generator
	PreparingStage *PreparingAttack // preparing data for attack
	CurrentAttackNumber int
	HTTPClient *http.Client
}

/*
preparing data for many request or one
1 map[string]interface return - params which insert in body request
2 map[string]string return - params which insert in query params of request
*/
func (atck *Attacker) preparingData(amountRequest int) *PreparingAttack {
	return nil
}

func (atck *Attacker) Reset() {
	atck.CurrentAttackNumber = 0
}

func NewAttacker(task *TaskBomber, generators generators.Generator) *Attacker{
	return &Attacker{
		Task: task,
		Generators: generators,
		PreparingStage: NewPreparingAttack(),
		CurrentAttackNumber: 0,
		HTTPClient: http.Client{
			Timeout: task.Task.CurrentTask.Scenario.Configuration
		}
	}
}

func (atck *Attacker) attack() error {
	stage := atck.Task.GetCurrentStage()
	request, err := http.NewRequest()
}
