package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"rest-bomber/models"

	"github.com/gorilla/mux"
)

type Application struct {
	conf   models.Configuration
	router *mux.Router
}

var app Application

//Parse settings from file argument path to this settings from command line
func (app *Application) parseSettings(path string) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := models.Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error Parse:", err)
	}
	app.conf = configuration
}

//Get path from command line arguments
func (app *Application) getPath() string {
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
