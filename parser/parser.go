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
	Reader  Reader
}

const (
	HEADERS = "headers"
	BODY    = "body"
)

//Parse - parsing data from file path and preparing this object to ParserPayload
func (parser *Parser) Parse(path string) {
	res := parser.Reader.read(path)

	var payload ParserPayload
	// var result map[string]interface{}

	if err := json.Unmarshal([]byte(res), &payload); err != nil {
		log.Println("Error when unmarshal json scenario", err)
	}
	parser.Payload = &payload
}
