package application

import (
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/models"
	"github.com/gorilla/mux"
)

type Application struct {
	Core            *core.Core
	Router          *mux.Router
	BomberCharacter *models.BomberCharacter
	GlobalConfig    *configuration.GlobalConfig
}

func NewApplication() *Application {
	return &Application{
		Core:            core.InitCore(),
		Router:          mux.NewRouter(),
		BomberCharacter: models.InitializeBomber(),
		GlobalConfig: configuration.,
	}
}

func newRouter() *mux.Router {
	return mux.NewRouter()
}
