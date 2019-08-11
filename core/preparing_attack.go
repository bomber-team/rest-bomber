package core

/*PreparingAttack - part of preparing data models for attack*/
type PreparingAttack struct {
	Core        *Core
	BodyParams  []map[string]interface{}
	QueryParams []map[string]string
}

/*GenerateSomeData - generate or use some value data from Input Schemas and scenario*/
func (prep *PreparingAttack) GenerateSomeData() {
	stage := prep.Core.State.GetCurrentStage()
	if prep.Core.State.SchemeOfCurrentStageEmpty() {
		return
	}

}
