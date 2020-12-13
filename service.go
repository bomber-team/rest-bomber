package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/handlers"
	"github.com/bomber-team/rest-bomber/helping"
	"github.com/sirupsen/logrus"
)

func workOnServiceSingal(signal int) {
	switch signal {
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

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	runtime.GOMAXPROCS(runtime.NumCPU())

	core := core.NewCore()
	coreHandler, errorHandling := handlers.NewCoreHandlers(core)
	if errorHandling != nil {
		logrus.Panic("Can not initialize consuming handler")
	}

	core.InitializeService()

	signalService := make(chan int)

	if err := coreHandler.InitTopicsHandlers(signalService); err != nil {
		signalService <- helping.FATALERROR
	}
	logrus.Info("Completed running service")
	sigOs := make(chan os.Signal, 1)
	signal.Notify(sigOs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case servSig := <-signalService:
		coreHandler.ShutdownToServer()
		workOnServiceSingal(servSig)
	case osSig := <-sigOs:
		coreHandler.ShutdownToServer()
		logrus.Error("Service catch signal terminated: ", osSig)
		os.Exit(1)
	}
}
