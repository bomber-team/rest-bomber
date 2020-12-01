package nats_listener

import (
	"github.com/google/uuid"
	"github.com/goreflect/gostructor"
	"github.com/sirupsen/logrus"
)

type NatsConnectionConfiguration struct {
	URL              string `cf_env:"NATS_URL" cf_default:"nats://localhost:4222"`
	NameClient       string `cf_env:"NATS_NAME" cf_default:"bomber"`
	MaxWait          int    `cf_env:"NATS_MAX_WAIT" cf_default:"1"`
	ReconnectDelay   int64  `cf_env:"NATS_RECONNECT_DELAY" cf_default:"2"`
	CurrentServiceID string `cf_env:"BOMBER_ID" cf_default:"15123kjnsjhad"`
}

func ParseConfiguration() (*NatsConnectionConfiguration, error) {
	parsed, errorConfiguration := gostructor.ConfigureSmart(&NatsConnectionConfiguration{}, "")
	if errorConfiguration != nil {
		logrus.Error(errorConfiguration)
		return nil, errorConfiguration
	}
	return parsed.(*NatsConnectionConfiguration), nil
}

func (config *NatsConnectionConfiguration) CorrectedGeneratingHandlerName() {
	uid, err := uuid.NewRandom()
	if err != nil {
		logrus.Error("Can not generating new handler name: ", err)
		return
	}
	config.CurrentServiceID = uid.String()
}
