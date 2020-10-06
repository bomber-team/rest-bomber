package handlers

import (
	"github.com/bomber-team/rest-bomber/core"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type CoreHandlers struct {
	connection      *nats.Conn
	currentHandlers []IHandlerTopic
}

func NewCoreHandlers(connection *nats.Conn, core *core.Core) (*CoreHandlers, error) {
	return &CoreHandlers{
		connection: connection,
		currentHandlers: []IHandlerTopic{
			newTopicHandler(connection, core),
		},
	}, nil
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
