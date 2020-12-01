package nats_listener

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	Connection *nats.Conn
	topic      string
}

func NewSubscriber(connection *nats.Conn, topicName string) *Subscriber {
	return &Subscriber{
		Connection: connection,
		topic:      topicName,
	}
}

func (subscr *Subscriber) Subscribe(handler nats.MsgHandler) error {
	subscription, err := subscr.Connection.Subscribe(subscr.topic, handler)
	if err != nil {
		return err
	}

	logrus.Info("Completed subscription: ", subscription)
	return nil
}
