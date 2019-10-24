package application

import (
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/routes"
	"github.com/gorilla/mux"
)

type Application struct {
	Core   *core.Core
	Router *mux.Router
}

func NewApplication() *Application {
	return &Application{
		Core:   core.InitCore(),
		Router: mux.NewRouter(),
	}
}

func newRouter() *mux.Router {
	routes := &routes.Routes{}
	return mux.NewRouter()
}
