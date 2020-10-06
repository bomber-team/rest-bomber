package handlers

type TestTopicHandler struct {
	subscriber *nats_listener.Subscriber
}

const (
	testTopicName = "test.1"
)

func newTopicHandler(conn *nats.Conn) *TestTopicHandler {
	return &TestTopicHandler{
		subscriber: nats_listener.NewSubscriber(testTopicName)
	}
}

func (handl *TestTopicHandler) Configuration() error {
	return handl.Subscribe(handl.handle)
}

func (handl *TestTopicHandler) handle(message *nats.Msg) {
	logrus.Info("Handled request by test topic handler. Subject: ", message.Subject, "Data: ", string(message.Data))
}
