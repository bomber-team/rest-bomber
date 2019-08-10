package routes

import (
	"encoding/json"
	"net/http"
	"rest-bomber/enhancer"
	"rest-bomber/payloads"

	"github.com/gorilla/mux"
	"gitlab.com/truecord_team/common/contents"
)

/*ScenarioRoute - route for scenarious*/
type ScenarioRoute struct {
	EResponser *enhancer.Responser
	Core       *core.Core
}

const (
	scenario = "/scenario"
)

func (router *ScenarioRoute) create(w http.ResponseWriter, request *http.Request) {
	var payload *payloads.ScenarioPayload
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		router.EResponser.ResponseWithError(w, request, http.StatusBadRequest, map[string]string{
			"status":    "error",
			"context":   "ScenarioRouter",
			"errorCode": err.Error(),
		}, contents.JSON)
	}
	// send to core scenarious
}

func (router *ScenarioRoute) remove(w http.ResponseWriter, request *http.Request) {
}

/*SettingRouter - setting routes by handlers*/
func (router *ScenarioRoute) SettingRouter(rout *mux.Router) *mux.Router {
	rout.HandleFunc(scenario, router.create).Methods("POST")
	rout.HandleFunc(scenario, router.remove).Methods("DELETE")
	return rout
}
