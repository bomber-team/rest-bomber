package routes

import (
	"github.com/bomber-team/rest-bomber/enhancer"
	"github.com/gorilla/mux"
)

type (
	/*ApplicationRouter - container for main routes service*/
	ApplicationRouter struct {
		Routes []IRoute
	}

	/*IRoute - interface for route*/
	IRoute interface {
		SettingRouter(router *mux.Router) *mux.Router
	}
)

/*
NewRoutes - setup inital routes
*/
func NewRoutes() *ApplicationRouter {
	responser := &enhancer.Responser{
		CurrentContentTypeResponse: enhancer.JSON,
	}
	return &ApplicationRouter{
		Routes: []IRoute{
			NewConfigureRoute(responser),
			newTaskRoute(responser),
		},
	}
}

/*ConfigureRoutes - configurating routes*/
func (setup *ApplicationRouter) ConfigureRoutes() *mux.Router {
	router := mux.NewRouter()
	for _, val := range setup.Routes {
		router = val.SettingRouter(router)
	}
	return router
}
