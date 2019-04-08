package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"os"
)

type Application struct {
	conf   Configuration
	router *mux.Router
}

//Configuration for attacker
type Configuration struct {
	AttackAddress      string `json:"attack-address"`
	ServerAddress      string `json:"server-address"`
	BufferSize         int    `json:"buffer-size"`
	Type               string `json:"type"`
	TimeBetweenAttacks string `json:"time-between-attacks"`
	AmountAttacks      int    `json:"amount-attacks"`
	PathScenarioFile   string `json:"path-scenario-file"`
	LogLevel           string `json:"log_level"`
	ApplicationType    string `json:"application-type"`
}

var app Application

//Parse settings from file argument path to this settings from command line
func (app *Application) parseSettings(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error Parse:", err)
	}
	app.conf = configuration
}

//Get path from command line arguments
func (app Application) getPath() string {
	path := flag.String("conf", "./settings.json", "Path to file with settings")
	flag.Parse()
	return *path
}

func init() {
	path := app.getPath()
	app.parseSettings(path)
}

func main() {

}
