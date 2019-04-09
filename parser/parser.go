package parser

import (
	"encoding/json"
	"log"
)

type Parser struct {
	Headers map[string]interface{}
	Body    map[string]interface{}
	Reader  Reader
}

const (
	HEADERS = "headers"
	BODY    = "body"
)

//Parse json
func (parser *Parser) Parse(path string) {
	res := parser.Reader.read(path)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(res), &result)

	if err != nil {
		log.Println("Error when unmarshal json scenario", err)
	}

	parser.Headers = result[HEADERS].(map[string]interface{})
	parser.Body = result[BODY].(map[string]interface{})
}
