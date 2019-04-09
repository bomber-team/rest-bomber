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
func (attacker Attacker) attack(conf models.Configuration) {
	attacker.parser.Parse(conf.PathScenarioFile)

	for i := 0; i < conf.AmountAttacks; i++ {
		buffer := attacker.generateBufferSize(conf.BufferSize)
		for j := 0; j < conf.BufferSize; j++ {
			go attacker.madeAttack(conf, buffer[j])
		}
	}
}

//Attack in goroutine
func (attacker Attacker) madeAttack(conf models.Configuration, body map[string]interface{}) {
	encodeBody, _ := json.Marshal(body)
	req, err := http.NewRequest(attacker.parser.Headers["method"].(string),
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
func (attacker Attacker) generateBufferSize(bufferSize int) (result []map[string]interface{}) {
	for j := 0; j < bufferSize; j++ {
		result = append(result, attacker.generator.GenerateForJsonBody(attacker.parser.Body))
	}
	return
}
