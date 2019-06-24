package routes

import (
	"encoding/json"
	"net/http"
	"rest-bomber/enhancer"
	"rest-bomber/payloads"

	"github.com/gorilla/mux"
)

/*InitRoute - */
type InitRoute struct {
	Responser *enhancer.Responser
	Routine   chan<- int
}

const (
	initRoute = "/init"
)

/*startWork -*/
func (route *InitRoute) startWork(writer http.ResponseWriter, request *http.Request) {
	var modelInit payloads.InitPayload
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&modelInit); err != nil {
		route.Responser.ResponseWithError(writer, request, http.StatusBadRequest, map[string]string{
			"status":  "error",
			"context": "initRoute",
			"code":    err.Error(),
		}, enhancer.JSON)
		return
	}

	switch modelInit.Field {
	case payloads.ActionStartWork:
		route.Routine <- payloads.ActionStartWorkID
		break
	case payloads.ActionStopWork:
		route.Routine <- payloads.ActionStopWorkID
		break
	default:
		route.Responser.ResponseWithError(writer, request, http.StatusBadRequest, map[string]string{
			"status":  "error",
			"context": "initRoute",
			"code":    "does not recognize, what i should do",
		}, enhancer.JSON)
		return
	}

	route.Responser.ResponseWithJSON(writer, request, http.StatusOK, map[string]string{
		"status": "ok",
		"code":   "work was accepted",
	}, enhancer.JSON)
	return
}

/*Configure - */
func (route *InitRoute) Configure(router *mux.Router) *mux.Router {
	router.HandleFunc(initRoute, route.startWork).Methods("POST")
	return router
}
