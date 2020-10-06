package nats_listener

import (
	"github.com/goreflect/gostructor"
	"github.com/sirupsen/logrus"
)

type NatsConnectionConfiguration struct {
	URL            string `cf_env:"NATS_URL" cf_default:"nats://localhost:4222"`
	NameClient     string `cf_env:"NATS_NAME" cf_default:"bomber"`
	MaxWait        int    `cf_env:"NATS_MAX_WAIT" cf_default:"1"`
	ReconnectDelay int64  `cf_env:"NATS_RECONNECT_DELAY" cf_default:"2"`
}

func ParseConfiguration() (*NatsConnectionConfiguration, error) {
	parsed, errorConfiguration := gostructor.ConfigureSmart(&NatsConnectionConfiguration{}, "")
	if errorConfiguration != nil {
		logrus.Error(errorConfiguration)
		return nil, errorConfiguration
	}
	return parsed.(*NatsConnectionConfiguration), nil
}
