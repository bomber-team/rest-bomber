package application

import (
	"github.com/bomber-team/rest-bomber/core"
	"github.com/gorilla/mux"
)

type Application struct {
	Core   *core.Core
	Router *mux.Router
}

func newApplication() *Application {
	return &Application{}
}
