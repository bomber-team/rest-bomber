package handlers

import (
	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type TaskTopicHandler struct {
	subscriber *nats_listener.Subscriber
	core       *core.Core
	bracket    chan int
}

const (
	taskTopicName = "bombers.tasks"
)

func newTaskTopicHandler(conn *nats.Conn, core *core.Core, config *nats_listener.NatsConnectionConfiguration) *TaskTopicHandler {
	return &TaskTopicHandler{
		subscriber: nats_listener.NewSubscriber(conn, testTopicName+config.CurrentServiceID),
		core:       core,
	}
}

func (handl *TaskTopicHandler) Configuration(signal chan int) error {
	errSubscription := handl.subscriber.Subscribe(handl.handle)
	handl.bracket = signal
	if errSubscription != nil {
		return errSubscription
	}
	return nil
}

func (handl *TaskTopicHandler) handle(message *nats.Msg) {
	logrus.Info("Handled request by task topic handler. Subject: ", message.Subject, "Data: ", string(message.Data))
	// starting task
	var paylaod rest_contracts.Task
	if err := paylaod.Unmarshal(message.Data); err != nil {
		logrus.Error("Can not unmarshal message from bomber server: ", err)
		return
	}

	logrus.Info("Starting working on task ID: ", paylaod.FormId)
}
