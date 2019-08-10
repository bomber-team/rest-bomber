package discovery

import (
	"net/http"
	"rest-bomber/enhancer"

	"github.com/gorilla/mux"
	"gitlab.com/truecord_team/common/contents"
)

/*Discovery - route for discovering by main backend service*/
type Discovery struct {
	EResponser          *enhancer.Responser
	Core                *core.Core
	IdentificatorBomber string
}

const (
	echo = "/"
)

/*echo - send to calling service id of our bomber*/
func (route *Discovery) echo(w http.ResponseWriter, request *http.Request) {
	route.EResponser.ResponseWithJSON(w, request, http.StatusOK, map[string]string{
		"stage": route.Core.State.CurrentStage,
		"id":    route.IdentificatorBomber,
	}, contents.JSON)
}

/*ConfigureDiscoveryRoute - configurating discovery route*/
func (route *Discovery) ConfigureDiscoveryRoute(router *mux.Router) *mux.Router {
	router.HandleFunc(echo, route.echo).Methods("GET")
	return router
}
