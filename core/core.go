package core

import (
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
)

type Core struct {
	Publisher *nats_listener.Publisher
}

func NewCore(conn *nats.Conn) *Core {
	return &Core{
		Publisher: nats_listener.NewPublisher(conn),
	}
}

func (core *Core) PreparingData( /*здесь контракт от сервака*/ ) /*здесь тоже какой-то контракт*/ {
	// готовим данные
}

func (core *Core) StartAttack() error {
	// запускаем атаку на какой-то маршрут
	return nil
}
