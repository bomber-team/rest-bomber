package discovery

import (
	"log"
	"net/http"
	"rest-bomber/configuration"
	"rest-bomber/core"
	"rest-bomber/enhancer"
	"strconv"

	"github.com/gorilla/mux"
)

type Discovery struct {
	GlobalConfig *configuration.GlobalConfiguration // global config, which contain bomber config
	Router       *mux.Router
}

/*createNewRouter - create new router for discovery*/
func (dscv *Discovery) createNewRouter(core *core.Core) {
	router := mux.NewRouter()
	route := DiscoveryRoute{
		EResponser:          &enhancer.Responser{},
		Core:                core,
		IdentificatorBomber: "",
	}
	router = route.ConfigureDiscoveryRoute(router)
	dscv.Router = router
}

/*InitHTTPListenerOnPort - create new goroutine for http listener*/
func (dscv *Discovery) InitHTTPListenerOnPort() {
	go func(dscv *Discovery) {
		if err := http.ListenAndServe(":"+strconv.Itoa(dscv.GlobalConfig.BomberConfigurationDiscovery.Port), dscv.Router); err != nil {
			log.Println("Error starting discovery listener")
		}
	}(dscv)
}
