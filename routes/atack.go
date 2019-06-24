package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-bomber/enhancer"
	"rest-bomber/payloads"
	"strings"

	"github.com/gorilla/mux"
)

/*AttackRoute - */
type AttackRoute struct {
	Responser *enhancer.Responser
	Routine   chan<- int
}

const (
	attackRoute = "/attack"
)

/*attackRoute - */
func (route *AttackRoute) attackRoute(writer http.ResponseWriter, request *http.Request) {
	log.Println("start attack route")
	var attackPayload payloads.AttackPayload
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&attackPayload); err != nil {
		route.Responser.ResponseWithError(writer, request, http.StatusBadRequest, map[string]string{
			"status":  "error",
			"context": "AttackRoute",
			"code":    err.Error(),
		}, enhancer.JSON)
		return
	}

	if strings.Compare(attackPayload.Field, payloads.StartAttack) == 0 {
		route.Routine <- 1
		route.Responser.ResponseWithJSON(writer, request, http.StatusOK, map[string]string{
			"status": "ok",
		}, enhancer.JSON)
		return
	}
	route.Responser.ResponseWithError(writer, request, http.StatusBadRequest, map[string]string{
		"status":  "error",
		"context": "attackRoute",
		"code":    "not handled need action",
	}, enhancer.JSON)
	return
}

/*InitRoute - */
func (route *AttackRoute) InitRoute(router *mux.Router) *mux.Router {
	router.HandleFunc(attackRoute, route.attackRoute).Methods("POST")
	return router
}
