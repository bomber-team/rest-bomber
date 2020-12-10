package handlers

import (
	"errors"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/bomber-team/rest-bomber/tools"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type CoreHandlers struct {
	connection      *nats.Conn
	currentHandlers []IHandlerTopic
	config          *nats_listener.NatsConnectionConfiguration
}

func NewCoreHandlers(core *core.Core) (*CoreHandlers, error) {

	return &CoreHandlers{
		connection: core.GetConnection(),
		currentHandlers: []IHandlerTopic{
			newTaskTopicHandler(core.GetConnection(), core, core.GetConfig()),
			newStarterTaskTopicHandler(core.GetConnection(), core, core.GetConfig()),
		},
		config: core.GetConfig(),
	}, nil
}

const (
	bomberInitTopic = "bombers.server.init_bomber"
	bomberDownTopic = "bombers.server.delete"
)

func (core *CoreHandlers) TestSendTask(config *nats_listener.NatsConnectionConfiguration) error {
	data := rest_contracts.Task{
		FormId: "test",
		Schema: &rest_contracts.RestSchema{
			Id:            "string",
			PathVariables: map[string]string{},
			Headers:       map[string]string{},
			Request:       []*rest_contracts.RequestParam{},
			Body:          []*rest_contracts.BodyParam{},
		},
		Script: &rest_contracts.RestScript{
			Address:       "http://127.0.0.1:8080",
			RequestMethod: "GET",
			Config: &rest_contracts.ConfigurationScript{
				Rps:  100000,
				Time: 50,
			},
		},
	}
	res, _ := data.Marshal()

	err := nats_listener.NewPublisher(core.connection).PublishNewMessage("bombers.tasks."+config.CurrentServiceID, res)
	if err != nil {
		logrus.Error("Error while init bomber: ", err)
	}
	return err
}

func (core *CoreHandlers) InitBomber() error {
	countTry := 0
	for {
		if countTry == 10 {
			return errors.New("Can not initialize bomber in server")
		}
		errInit := core.initBomber()
		if errInit != nil {
			countTry++
			continue
		}
		break
	}
	return nil
}

func (core *CoreHandlers) initBomber() error {
	res := rest_contracts.InitBomberPayload{
		BomberIp: tools.InitIp(),
		BomberId: core.config.CurrentServiceID,
	}
	data, _ := res.Marshal()

	err := nats_listener.NewPublisher(core.connection).PublishNewMessage(bomberInitTopic, data)
	if err != nil {
		logrus.Error("Error while init bomber: ", err)
	}
	return err
}

func (core *CoreHandlers) InitTopicsHandlers(signal chan int) error {
	logrus.Info("Starting configuring topic handlers")
	for _, handler := range core.currentHandlers {
		if err := handler.Configuration(signal); err != nil {
			return err
		}
	}
	logrus.Info("Completed configuring topic handlers")
	return nil
}

func (core *CoreHandlers) ShutdownToServer() {
	res := rest_contracts.InitBomberPayload{
		BomberIp: tools.InitIp(),
		BomberId: core.config.CurrentServiceID,
	}
	data, _ := res.Marshal()
	err := nats_listener.NewPublisher(core.connection).PublishNewMessage(bomberDownTopic, data)
	if err != nil {
		logrus.Error("Error while send deleting bomber from server: ", err)
	}
}
