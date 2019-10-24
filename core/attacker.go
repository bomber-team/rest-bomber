package core

/*Attacker - main entry for attack start*/
type Attacker struct {
	Task           *TaskBomber      // contain current schemas and scenario
	PreparingStage *PreparingAttack // preparing data for attack
}

/*
preparing data for many request or one
1 map[string]interface return - params which insert in body request
2 map[string]string return - params which insert in query params of request
*/
func (atck *Attacker) preparingData(amountRequest int) ([]map[string]interface{}, []map[string]string) {

	return nil, nil
}

func (atck *Attacker) attack() {

}
