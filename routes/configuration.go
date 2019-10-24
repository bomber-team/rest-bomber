package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/enhancer"
	"github.com/bomber-team/rest-bomber/routes/payloads"
	"github.com/gorilla/mux"
)

/*ConfigurationRoute - route for setting configuration*/
type ConfigurationRoute struct {
	EResponser *enhancer.Responser
	Core       *core.Core
}

const (
	configBomber = "/configurate"
)

func (router *ConfigurationRoute) configureBomber(w http.ResponseWriter, request *http.Request) {
	var payload *payloads.BomberConfig
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		router.EResponser.ResponseWithError(w, request, http.StatusBadRequest, map[string]string{
			"status":    "error",
			"context":   "BomberConfigurationRouter",
			"errorCode": err.Error(),
		}, enhancer.JSON)
	}
	// send to core new configuration for bomber
}

/*SettingRouter - setting routes by handlers*/
func (router *ConfigurationRoute) SettingRouter(rout *mux.Router) *mux.Router {
	rout.HandleFunc(configBomber, router.configureBomber).Methods("POST")
	return rout
}
