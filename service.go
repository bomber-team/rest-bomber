package main

import (
	"os"

	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/handlers"
	"github.com/bomber-team/rest-bomber/helping"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	parsedConfigureService, errParsing := nats_listener.ParseConfiguration()
	if errParsing != nil {
		logrus.Error("can not parsed configuration: ", errParsing)
		panic(errParsing)
	}
	parsedConfigureService.CorrectedGeneratingHandlerName()
	connection, errConnection := nats_listener.CreateNewConnectionToNats(parsedConfigureService)
	if errConnection != nil {
		logrus.Error("Can not connected to nats: ", errConnection)
		panic(errConnection)
	}
	core := core.NewCore(connection)
	coreHandler, errorHandling := handlers.NewCoreHandlers(connection, core, parsedConfigureService)
	if errorHandling != nil {
		logrus.Panic("Can not initialize consuming handler")
	}

	signalService := make(chan int)

	if err := coreHandler.InitTopicsHandlers(signalService); err != nil {
		signalService <- helping.FATALERROR
	}

	logrus.Info("Completed running service")

	switch <-signalService {
	case helping.STOPSERVICE:
		logrus.Info("Shutdown service completely")
		os.Exit(0)
	case helping.FATALERROR:
		logrus.Error("Error while working service with fatal errors")
		os.Exit(1)
	default:
		logrus.Error("Not recognized signal")
		os.Exit(1)
	}
}
