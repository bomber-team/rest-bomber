package core

/*Attacker - main entry for attack start*/
type Attacker struct {
	Core           *Core            // contain current schemas and scenario
	PreparingStage *PreparingAttack // preparing data for attack
	BodyMap        []map[string]interface{}
	QueryMap       []map[string]string
}

/*Pipeline - pipe for attack run*/
func (atck *Attacker) Pipeline() {

}

/*
preparing data for many request or one
1 map[string]interface return - params which insert in body request
2 map[string]string return - params which insert in query params of request
*/
func (atck *Attacker) preparingData(amountRequest int) ([]map[string]interface{}, []map[string]string) {
	atck.BodyMap = make([]map[string]interface{}, amountRequest)
	atck.QueryMap = make([]map[string]string, amountRequest)
	return nil, nil
}

func (atck *Attacker) attack() {

}
