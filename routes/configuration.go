package routes

import (
	"encoding/json"
	"net/http"
	"rest-bomber/enhancer"
	"rest-bomber/models"

	"github.com/gorilla/mux"
)

/*ConfigurationRoute - route for setting configuration*/
type ConfigurationRoute struct {
	Responser *enhancer.Responser
	Routine   chan<- int
}

const (
	routerForConfig = "/configurate"
)

/*configurationRoute - route for setting attacker to attack*/
func (route *ConfigurationRoute) configurationRoute(writer http.ResponseWriter, request *http.Request) {
	var modelConfiguration models.Configuration
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&modelConfiguration); err != nil {
		route.Responser.ResponseWithError(writer, request, http.StatusBadRequest, map[string]string{
			"status":  "error",
			"context": "configuration",
			"code":    err.Error(),
		}, enhancer.JSON)
		return
	}

	if modelConfiguration.CheckValid() {
		route.Responser.ResponseWithError(writer, request, http.StatusBadRequest, map[string]string{
			"status":  "error",
			"context": "configuration",
			"code":    "configuration model not valid",
		}, enhancer.JSON)
		return
	}

	route.Responser.ResponseWithJSON(writer, request, http.StatusOK, map[string]string{
		"status":  "ok",
		"context": "configuration",
	}, enhancer.JSON)
	return
}

/*InitRoute -*/
func (route *ConfigurationRoute) InitRoute(router *mux.Router) *mux.Router {
	router.HandleFunc(routerForConfig, route.configurationRoute).Methods("POST")
	return router
}
