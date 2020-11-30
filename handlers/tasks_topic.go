package handlers

import (
	"encoding/json"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type TaskTopicHandler struct {
	subscriber *nats_listener.Subscriber
	publisher  *nats_listener.Publisher
	core       *core.Core
	bracket    chan int
	config     *nats_listener.NatsConnectionConfiguration
}

const (
	taskTopicName     = "bombers.tasks."
	taskStatusChanger = "bombers.server.task_status"
)

const (
	CONFIGURED          = 0
	COMPLETED_ATTACK    = 1
	ERROR_CONFIGURATION = 3
	ERROR_ATTACK        = 2
)

type ResultConfiguration struct {
	TaskID string `json:"task_id"`
	Result int    `json:"status"`
}

func newTaskTopicHandler(conn *nats.Conn, core *core.Core, config *nats_listener.NatsConnectionConfiguration) *TaskTopicHandler {
	return &TaskTopicHandler{
		subscriber: nats_listener.NewSubscriber(conn, taskTopicName+config.CurrentServiceID),
		publisher:  nats_listener.NewPublisher(conn),
		core:       core,
		config:     config,
	}
}

func (handl *TaskTopicHandler) Configuration(signal chan int) error {
	logrus.Info("Start task topic handler")
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
	logrus.Info("Starting building ", paylaod.Script.Config.Rps*paylaod.Script.Config.Time, " amount request")
	handl.core.PreparingData(paylaod)
	logrus.Info("Completed builded Requests for attack")
	formatResultStatusTask(paylaod.FormId, CONFIGURED, handl.publisher)
	handl.publisher.PublishNewMessage(taskTopicStarter+handl.config.CurrentServiceID, message.Data)
}

func formatResultStatusTask(taskId string, status int, publisher *nats_listener.Publisher) {
	resultMarshaled, err := json.Marshal(ResultConfiguration{
		TaskID: taskId,
		Result: status,
	})
	if err != nil {
		logrus.Error("Error forming result for backend")
		return
	}
	errPublish := publisher.PublishNewMessage(taskStatusChanger, resultMarshaled)
	if errPublish != nil {
		logrus.Error("Error while publish status by task: ", errPublish)
	}
}
