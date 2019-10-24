package tests

import (
	parser2 "restbomber/parser"
	"testing"
)

func Test_ReadFromFile(t *testing.T) {
	parser := parser2.Parser{
		Reader: parser2.FileReader{},
	}

	parser.Parse("/home/kostya05983/Projects/restbomber/tests/tests_data/scenario.json")
}
