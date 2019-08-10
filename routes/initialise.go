package routes

import "github.com/gorilla/mux"

type (
	/*Routes - container for main routes service*/
	Routes struct {
		Routes []IRoute
	}

	/*IRoute - interface for route*/
	IRoute interface {
		SettingRouter(router *mux.Router) *mux.Router
	}
)

/*ConfigureRoutes - configurating routes*/
func (setup *Routes) ConfigureRoutes(router *mux.Router) *mux.Router {
	var routerUpdate *mux.Router
	for _, val := range setup.Routes {
		routerUpdate = val.SettingRouter(router)
	}
	return routerUpdate
}
