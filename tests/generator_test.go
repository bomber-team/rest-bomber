package tests

import (
	"rest-bomber/generators"
	parser2 "rest-bomber/parser"
	"testing"
)

//Test checks for nil when generates
func Test_generator(t *testing.T) {
	parser := parser2.Parser{
		Reader: parser2.FileReader{},
	}

	parser.Parse("../tests/tests_data/scenario.json")

	generator := generators.Generator{}
	var result = make(map[string]interface{})
	generator.GenerateForJsonBody(parser.Body, result)
}
