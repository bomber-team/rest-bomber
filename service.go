package main

import (
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {

	subscriber1 := nats_listener.NewSubscriber(connection, "bomber.tasks.1")
	subscriber1.Subscribe(func(message *nats.Msg) {
		logrus.Info("Handled new message from topic: ", message.Subject, " Message: ", string(message.Data))
	})
}
