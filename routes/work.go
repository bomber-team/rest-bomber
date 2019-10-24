package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WorkRoute struct {
	// EResponser *enhancer.Responser
}

const (
	work = "/work"
)

func (router *WorkRoute) workStage(w http.ResponseWriter, request *http.Request) {

}

/*SettingRouter - setting routes by handlers*/
func (router *WorkRoute) SettingRouter(rout *mux.Router) *mux.Router {
	rout.HandleFunc(work, router.workStage).Methods("POST")
	return rout
}
