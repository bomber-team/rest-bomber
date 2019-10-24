package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bomber-team/rest-bomber/enhancer"
	"github.com/bomber-team/rest-bomber/routes/payloads"
	"github.com/gorilla/mux"
)

type WorkRoute struct {
	EResponser *enhancer.Responser
	TaskChan   *chan int
}

const (
	work = "/work"
)

func (router *WorkRoute) ConfigureTaskChan(taskchan chan int) {
	router.TaskChan = &taskchan
}

func (router *WorkRoute) workStage(w http.ResponseWriter, request *http.Request) {
	var workModel payloads.WorkPayload
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&workModel); err != nil {
		router.EResponser.ResponseWithError(w, request, http.StatusBadRequest, map[string]string{
			"status":    "error",
			"context":   "WorkRoute",
			"errorCode": err.Error(),
		}, enhancer.JSON)
	}
	*(router.TaskChan) <- workModel.Action
	router.EResponser.ResponseWithError(w, request, http.StatusOK, map[string]string{
		"status":  "success",
		"context": "WorkRoute.workStage",
		"message": "action will setup",
	}, enhancer.JSON)
}

/*SettingRouter - setting routes by handlers*/
func (router *WorkRoute) SettingRouter(rout *mux.Router) *mux.Router {
	rout.HandleFunc(work, router.workStage).Methods("POST")
	return rout
}
