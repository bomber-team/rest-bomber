package configuration_discovery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*ConfigurationDiscovery - config for discovery route*/
type ConfigurationDiscovery struct {
	Port int `json:"port"`
}

//Read file from disk
//returns string for future parsing
func (conf *ConfigurationDiscovery) read(path string) []byte {
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

/*ParseFieldFile - file configuration reading*/
func (conf *ConfigurationDiscovery) ParseFieldFile(path string) error {
	bytes := conf.read(path)
	if err := json.Unmarshal(bytes, conf); err != nil {
		log.Println("Can not unmarshal data")
		return err
	}
	return nil
}
