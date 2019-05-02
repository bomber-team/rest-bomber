package parser

import (
	"encoding/json"
	"log"
)

/*ParserPayload - enter packet*/
type ParserPayload struct {
	Headers map[string]interface{} `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

/*Parser - structure for contain */
type Parser struct {
	Payload *ParserPayload
	Reader
}

const (
	HEADERS = "headers"
	BODY    = "body"
)

//Parse - parsing data from file path and preparing this object to ParserPayload
func (parser *Parser) Parse(path string) {
	res := parser.read(path)

	var payload ParserPayload

	if err := json.Unmarshal(res, &payload); err != nil {
		log.Println("Error when unmarshal json scenario", err)
	}
	parser.Payload = &payload
}
