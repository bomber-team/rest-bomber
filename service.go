package main

import (
	"net"
	"os"
	"runtime"

	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/handlers"
	"github.com/bomber-team/rest-bomber/helping"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	runtime.GOMAXPROCS(runtime.NumCPU())
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
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	result := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = ipnet.IP.To4().String()
			}
		}
	}
	core := core.NewCore(connection, result)
	coreHandler, errorHandling := handlers.NewCoreHandlers(connection, core, parsedConfigureService)
	if errorHandling != nil {
		logrus.Panic("Can not initialize consuming handler")
	}

	for {
		errInit := coreHandler.InitBomber(parsedConfigureService)
		if errInit != nil {
			continue
		}
		break
	}

	signalService := make(chan int)

	if err := coreHandler.InitTopicsHandlers(signalService); err != nil {
		signalService <- helping.FATALERROR
	}
	// coreHandler.TestSendTask(parsedConfigureService)
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
