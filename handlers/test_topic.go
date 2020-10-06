package handlers

import (
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type TestTopicHandler struct {
	subscriber *nats_listener.Subscriber
	core       *core.Core
	bracket    chan int
}

const (
	testTopicName = "test.1"
)

func newTopicHandler(conn *nats.Conn, core *core.Core) *TestTopicHandler {
	return &TestTopicHandler{
		subscriber: nats_listener.NewSubscriber(conn, testTopicName),
		core:       core,
	}
}

func (handl *TestTopicHandler) Configuration(signal chan int) error {

	errSubscription := handl.subscriber.Subscribe(handl.handle)
	if errSubscription != nil {
		return errSubscription
	}
	return nil
}

func (handl *TestTopicHandler) handle(message *nats.Msg) {
	logrus.Info("Handled request by test topic handler. Subject: ", message.Subject, "Data: ", string(message.Data))
}
