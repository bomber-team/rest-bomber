package nats_listener

type ISubscriber interface {
	Subscribe() error
}
