package discovery

import (
	"encoding/json"
	"net/http"
	"restbomber/configuration"
	"restbomber/core"
	"restbomber/enhancer"
	"restbomber/payloads"

	"github.com/gorilla/mux"
	"gitlab.com/truecord_team/common/contents"
)

/*DiscoveryRoute - route for discovering by main backend service*/
type DiscoveryRoute struct {
	EResponser          *enhancer.Responser
	Core                *core.Core
	GlobalConfig        *configuration.GlobalConfiguration // create new type, which contain config in main layer
	IdentificatorBomber string
}

const (
	echo = "/"
)

/*echo - send to calling service id of our bomber*/
func (route *DiscoveryRoute) echo(w http.ResponseWriter, request *http.Request) {
	route.EResponser.ResponseWithJSON(w, request, http.StatusOK, map[string]string{
		"id": route.IdentificatorBomber,
	}, contents.JSON)
}

func (route *DiscoveryRoute) configureBomber(w http.ResponseWriter, request *http.Request) {
	var payload *payloads.BomberConfig
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, request, http.StatusBadRequest, map[string]string{
			"status":  "error",
			"context": "discovery_route",
			"code":    err.Error(),
		}, contents.JSON)
	}
	route.GlobalConfig.SetupBomberConfiguration(payload)
	route.EResponser.ResponseWithJSON(w, request, http.StatusOK, map[string]string{
		"status":  "ok",
		"context": "discovery_route",
	}, contents.JSON)
}

/*ConfigureDiscoveryRoute - configurating discovery route*/
func (route *DiscoveryRoute) ConfigureDiscoveryRoute(router *mux.Router) *mux.Router {
	router.HandleFunc(echo, route.echo).Methods("GET")
	router.HandleFunc(echo, route.configureBomber).Methods("POST")
	return router
}
