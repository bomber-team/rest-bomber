package routes

import (
	"net/http"
	"rest-bomber/enhancer"

	"github.com/gorilla/mux"
)

type SchemeRoute struct {
	EResponser *enhancer.Responser
}

const (
	scheme = "/schemes"
)

func (router *SchemeRoute) create(w http.ResponseWriter, request *http.Request) {

}

func (router *SchemeRoute) remove(w http.ResponseWriter, request *http.Request) {

}

/*SettingRouter - setting routes by handlers*/
func (router *SchemeRoute) SettingRouter(rout *mux.Router) *mux.Router {
	rout.HandleFunc(scheme, router.create).Methods("POST")
	rout.HandleFunc(scheme, router.remove).Methods("DELETE")
	return rout
}
