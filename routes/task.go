package routes

import (
	"encoding/json"
	"net/http"

	"github.com/bomber-team/rest-bomber/enhancer"
	"github.com/bomber-team/rest-bomber/routes/payloads"
	"github.com/gorilla/mux"
)

type TaskRoute struct {
	EResponser *enhancer.Responser
	TaskChan   *chan payloads.Task
}

func newTaskRoute(responer *enhancer.Responser) *TaskRoute {
	return &TaskRoute{
		EResponser: responer,
	}
}

const (
	task = "/task"
)

func (router *TaskRoute) ConfigureTaskChann(taskchan chan payloads.Task) {
	router.TaskChan = &taskchan
}

func (router *TaskRoute) create(w http.ResponseWriter, request *http.Request) {
	var taskModel payloads.Task
	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&taskModel); err != nil {
		router.EResponser.ResponseWithError(w, request, http.StatusBadRequest, map[string]string{
			"status":    "error",
			"context":   "TaskRoute",
			"errorCode": err.Error(),
		}, enhancer.JSON)
	}
	*(router.TaskChan) <- taskModel
	router.EResponser.ResponseWithError(w, request, http.StatusOK, map[string]string{
		"status":  "success",
		"context": "TaskRoute.create",
		"message": "task will setup",
	}, enhancer.JSON)
}

/*SettingRouter - setting routes by handlers*/
func (router *TaskRoute) SettingRouter(rout *mux.Router) *mux.Router {
	rout.HandleFunc(task, router.create).Methods("POST")
	return rout
}
