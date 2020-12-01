package nats_listener

import "github.com/nats-io/nats.go"

type Publisher struct {
	Connection *nats.Conn
}

func NewPublisher(connection *nats.Conn) *Publisher {
	return &Publisher{
		Connection: connection,
	}
}

func (publsh *Publisher) PublishNewMessage(topic string, message []byte) error {
	return publsh.Connection.Publish(topic, message)
}
