package main

import (
	"log"
	"net/http"
	"rest-bomber/enhancer"
	"rest-bomber/payloads"
	"rest-bomber/routes"

	"github.com/gorilla/mux"
)

type Application struct {
	InitRoute          *routes.InitRoute
	ConfigurationRoute *routes.ConfigurationRoute
	AttackRoute        *routes.AttackRoute
	Routine            chan int
}

var router *mux.Router
var app Application
var Routine chan int

// //Parse settings from file argument path to this settings from command line
// func (app *Application) parseSettings(path string) {
// 	file, _ := os.Open(path)
// 	defer file.Close()
// 	decoder := json.NewDecoder(file)
// 	configuration := models.Configuration{}
// 	err := decoder.Decode(&configuration)
// 	if err != nil {
// 		fmt.Println("error Parse:", err)
// 	}
// 	app.conf = configuration
// }

// //Get path from command line arguments
// func (app *Application) getPath() string {
// 	path := flag.String("conf", "./settings.json", "Path to file with settings")
// 	flag.Parse()
// 	return *path
// }

func init() {
	// path := app.getPath()
	// app.parseSettings(path)
	// channelForJob := make(chan int, 1)
	Routine := make(chan int)
	responser := &enhancer.Responser{
		CurrentContentTypeResponse: enhancer.JSON,
	}
	app := &Application{
		InitRoute: &routes.InitRoute{
			Responser: responser,
			Routine:   Routine,
		},
		ConfigurationRoute: &routes.ConfigurationRoute{
			Responser: responser,
			Routine:   Routine,
		},
		AttackRoute: &routes.AttackRoute{
			Responser: responser,
			Routine:   Routine,
		},
		Routine: Routine,
	}

	rout := mux.NewRouter()
	rout = app.InitRoute.Configure(rout)
	rout = app.ConfigurationRoute.InitRoute(rout)
	rout = app.AttackRoute.InitRoute(rout)
	router = rout
}

func handlingRoutineActions() {
	log.Print("worker for device starting")
	workJob int <- Routine
	log.Print("in channel handle new value")
	switch workJob {
	case payloads.ActionStartWorkID:
		log.Print("start work")
		break
	case payloads.ActionStopWorkID:
		log.Print("stop work")
		break
	case payloads.ActionWriteConfigurationID:
		log.Print("start configration build")
		break
	case payloads.ActionWriteScenarioID:
		log.Print("start scenario build")
		break
	case payloads.ActionWriteSchemeID:
		log.Print("start scheme build")
		break
	default:
		continue
	}

}

func main() {
	go handlingRoutineActions()
	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("start test request")
		Routine <- payloads.ActionStartWorkID
		log.Print("success send new routine for worker")

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(""))
	}).Methods("GET")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf(err.Error())
	}
}
