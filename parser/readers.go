package parser

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Reader interface {
	read(path string) string
}

type FileReader struct {
}

//Read file from disk
//returns string for future parsing
func (reader *FileReader) read(path string) string {
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully open" + path)
	// defer the closing of our jsonFile so that we can Parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return string(byteValue)
}

type ApiReader struct{}

//Read scenario from api, path address of server
//returns string for future parsing
func (reader *ApiReader) read(path string) string {
	return ""
}
