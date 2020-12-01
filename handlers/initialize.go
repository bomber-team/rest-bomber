package handlers

import (
	"net"
	"os"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type CoreHandlers struct {
	connection      *nats.Conn
	currentHandlers []IHandlerTopic
	config          *nats_listener.NatsConnectionConfiguration
}

func NewCoreHandlers(connection *nats.Conn, core *core.Core, config *nats_listener.NatsConnectionConfiguration) (*CoreHandlers, error) {
	return &CoreHandlers{
		connection: connection,
		currentHandlers: []IHandlerTopic{
			newTaskTopicHandler(connection, core, config),
			newStarterTaskTopicHandler(connection, core, config),
		},
	}, nil
}

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
				Rps:  1000,
				Time: 10,
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

func (core *CoreHandlers) InitBomber(config *nats_listener.NatsConnectionConfiguration) error {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	result := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = ipnet.IP.To4().String()
			}
		}
	}
	res := rest_contracts.InitBomberPayload{
		BomberIp: result,
		BomberId: config.CurrentServiceID,
	}
	data, _ := res.Marshal()

	err = nats_listener.NewPublisher(core.connection).PublishNewMessage("bombers.server.init_bomber", data)
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
