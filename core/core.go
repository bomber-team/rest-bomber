package core

import (
	"time"

	"github.com/bomber-team/bomber-proto-contracts/golang/system"
	"github.com/bomber-team/rest-bomber/nats_listener"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type Core struct {
	publisher           *nats_listener.Publisher
	config              *nats_listener.NatsConnectionConfiguration
	currentStatusBomber system.StatusBomber
}

const (
	bomberStatus = "bomber.status"
	bomberResult = "bomber.result"
)

func NewCore(conn *nats.Conn) *Core {
	return &Core{
		publisher:           nats_listener.NewPublisher(conn),
		currentStatusBomber: system.StatusBomber_UP,
	}
}

func (core *Core) PreparingData( /*здесь контракт от сервака*/ ) /*здесь тоже какой-то контракт*/ {
	// готовим данные
}

func (core *Core) StartAttack() error {
	// запускаем атаку на какой-то маршрут
	core.currentStatusBomber = system.StatusBomber_WORKING
	return nil
}

func (core *Core) InitializeService() {
	core.changeStatusBomber(core.currentStatusBomber)
}

func (core *Core) handlingChangeStatusBomber() {
	currentStatus := core.currentStatusBomber
	for {
		time.Sleep(time.Second * 5)
		if currentStatus != core.currentStatusBomber {
			logrus.Info("Handled changing current status worker: ", core.currentStatusBomber.String())
			core.changeStatusBomber(core.currentStatusBomber)
			currentStatus = core.currentStatusBomber
		}
	}
}

func (core *Core) gracefullDownService() {
	logrus.Info("Graceful down service")
	core.changeStatusBomber(system.StatusBomber_DOWN)
}

func (core *Core) changeStatusBomber(status system.StatusBomber) {
	statusBomberInitialized := system.BomberStatusChange{
		BomberId:     core.config.CurrentServiceID,
		StatusBomber: status,
	}
	data, errMarshaling := statusBomberInitialized.Marshal()
	if errMarshaling != nil {
		logrus.Error("Can not marshaled payload for bomber server: ", errMarshaling)
	}
	if errPublish := core.publisher.PublishNewMessage(bomberStatus, data); errPublish != nil {
		logrus.Error("Can not publish message into broker nats")
	}
}
