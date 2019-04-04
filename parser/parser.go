package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Parser struct {
}

func (parser Parser) parse(file string) map[string]interface{} {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully open" + file)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
}
