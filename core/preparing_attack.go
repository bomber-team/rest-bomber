package core

import "github.com/bomber-team/rest-bomber/models"

/*PreparingAttack - part of preparing data models for attack*/
type PreparingAttack struct {
	CurrentTask *TaskBomber
	// Generator
	BodyParams  []map[string]interface{}
	QueryParams []map[string]string
}

func NewPreparingAttack(task *TaskBomber) *PreparingAttack{
	return &PreparingAttack{
		CurrentTask: task,
	}
}

/*GenerateSomeData - generate or use some value data from Input Schemas and scenario*/
func (prep *PreparingAttack) GenerateSomeData() {
	stage := prep.CurrentTask.GetCurrentStage()
	currentScheme := prep.CurrentTask.GetScheme(stage.Scheme)
	prep.BodyParams = make([]map[string]interface{}, prep.CurrentTask.CurrentTask.Scenario.Configuration.AmountRequests)

	for i:=0; i < prep.CurrentTask.CurrentTask.Scenario.Configuration.AmountRequests; i++ {
		prep.GenerateQuery(currentScheme)
		prep.GenerateBody(currentScheme)
	}
}

func (prep *PreparingAttack) GenerateQuery(currentIndex int, scheme *models.Scheme) {
	currentQueryParam := scheme.Params.Query[currentIndex]
	value := ""
	if currentQueryParam.GeneratorNeed {
		//execute Generator
	}
	prep.QueryParams[currentIndex][currentQueryParam.Name] = value
}

func (prep *PreparingAttack) GenerateBody(currentIndex int, scheme *models.Scheme) {
	currentBodyParam :=scheme.Params.Body[currentIndex]
	var value interface{}

}
