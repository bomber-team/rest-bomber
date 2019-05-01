package attacker

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"rest-bomber/generators"
	"rest-bomber/models"
	parser2 "rest-bomber/parser"
)

type Attacker struct {
	parser    parser2.Parser
	generator generators.Generator
}

//Realize algorithm of attack on server
func (atck *Attacker) attack(conf models.Configuration) {
	atck.parser.Parse(conf.PathScenarioFile)

	for i := 0; i < conf.AmountAttacks; i++ {
		buffer := atck.generateBufferSize(conf.BufferSize)
		for j := 0; j < conf.BufferSize; j++ {
			go atck.madeAttack(conf, buffer[j])
		}
	}
}

//Attack in goroutine
func (atck *Attacker) madeAttack(conf models.Configuration, body map[string]interface{}) {
	encodeBody, _ := json.Marshal(body)
	req, err := http.NewRequest(atck.parser.Headers["method"].(string),
		conf.ServerAddress, bytes.NewBuffer(encodeBody))
	if err != nil {
		println("Can't made request")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error while wait request")
	}
	defer resp.Body.Close()
	log.Println("Get response")
}

//Generate our buffer
func (atck *Attacker) generateBufferSize(bufferSize int) (result []map[string]interface{}) {
	for j := 0; j < bufferSize; j++ {
		result = append(result, atck.generator.GenerateForJSONBody(atck.parser.Body))
	}
	return
}
