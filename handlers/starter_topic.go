package handlers

import (
	"sync"
	"time"

	"github.com/bomber-team/bomber-proto-contracts/golang/rest_contracts"
	"github.com/bomber-team/rest-bomber/core"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type StarterTopicHandler struct {
	subscriber *nats_listener.Subscriber
	publisher  *nats_listener.Publisher
	core       *core.Core
	bracket    chan int
}

const (
	taskTopicStarter = "bombers.starter.tasks."
	taskTopicResult  = "bombers.server.task_result"
	taskStatusResult = "bombers.server.task_status"
)

func newStarterTaskTopicHandler(conn *nats.Conn, core *core.Core, config *nats_listener.NatsConnectionConfiguration) *StarterTopicHandler {
	return &StarterTopicHandler{
		subscriber: nats_listener.NewSubscriber(conn, taskTopicStarter+config.CurrentServiceID),
		publisher:  nats_listener.NewPublisher(conn),
		core:       core,
	}
}

func (handl *StarterTopicHandler) Configuration(signal chan int) error {
	logrus.Info("Start starter topic handler")
	errSubscription := handl.subscriber.Subscribe(handl.handle)
	handl.bracket = signal
	if errSubscription != nil {
		return errSubscription
	}
	return nil
}

func (handl *StarterTopicHandler) handle(message *nats.Msg) {
	logrus.Info("Handled request by task topic starter handler. Subject: ", message.Subject, "Data: ", string(message.Data))
	// starting task
	var paylaod rest_contracts.Task
	if err := paylaod.Unmarshal(message.Data); err != nil {
		logrus.Error("Can not unmarshal message from bomber server: ", err)
		return
	}

	logrus.Info("Starting checking reading for attack: ", paylaod.FormId)
	if handl.core.CheckReady() {
		logrus.Debug("Attack REady")
		var wg sync.WaitGroup
		wg.Add(1)
		timeStart := time.Now()
		handl.core.Start(paylaod, &wg)
		wg.Wait()
		timeEnd := time.Since(timeStart)
		logrus.Debug("Attacks completed. Start extracting data")
		result := handl.core.FormResultAttack()
		result.ElapsedTimeAttack = timeEnd.Nanoseconds()
		result.BomberId = handl.core.GetConfig().CurrentServiceID
		logrus.Debug("Summary estimated time for attack: ", timeEnd.Nanoseconds(), " ns")
		marshaledData, err := result.Marshal()
		if err != nil {
			logrus.Error("Error marshaled result attack: ", err)
			formatResultStatusTask(paylaod.FormId, ERROR_ATTACK, handl.publisher)
			return
		}
		handl.publisher.PublishNewMessage(taskTopicResult, marshaledData)
		formatResultStatusTask(paylaod.FormId, COMPLETED_ATTACK, handl.publisher)
	}
}
