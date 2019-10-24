package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/bomber-team/rest-bomber/application"
	"github.com/bomber-team/rest-bomber/discovery"
	"github.com/bomber-team/rest-bomber/models"
)

var app *application.Application

func init() {
	app = application.NewApplication()
	bomba := models.InitializeBomber()
	speaker := discovery.Speaker{
		Timeout:         1,
		Address:         "224.0.0.1:9999",
		MaxDatagramSize: 8192,
		Packet:          bomba.GetBytes(),
	}
	go speaker.RunEcho()
}

func readFileJson(filename string) []byte {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func main() {
	for {
		fmt.Println("test")
		time.Sleep(time.Hour)
	}
}
