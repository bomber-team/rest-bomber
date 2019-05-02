package parser

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*Reader - main interface for read datas from other sources*/
type Reader interface {
	read(path string) []byte
}

/*FileReader - main structure for read files from local disk schemes*/
type FileReader struct {
}

//Read file from disk
//returns string for future parsing
func (reader FileReader) read(path string) []byte {
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully open" + path)
	// defer the closing of our jsonFile so that we can Parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

type ApiReader struct{}

//Read scenario from api, path address of server
//returns string for future parsing
func (reader *ApiReader) read(path string) string {
	return ""
}
