package handlers

type CoreHandlers struct {
	connection *nats.Conn
	currentHandlers []ITopicHandler
}

func NewCoreHandlers() (*CoreHandlers, error) {
	parsedConfigureService, errParsing := nats_listener.ParseConfiguration()
	if errParsing != nil {
		logrus.Error("can not parsed configuration: ", errParsing)
		return nil, errParsing
	}
	connection, errConnection := nats_listener.CreateNewConnectionToNats(parsedConfigureService)
	if errConnection != nil {
		logrus.Error("Can not connected to nats: ", errConnection)
		return nil, errConnection
	}
	return &CoreHandlers{
		connection: connection,
		currentHandlers: []ITopicHandler {
			newTopicHandler(connection)
		}
	}, nil
}

func (core *CoreHandlers) InitTopicsHandlers() error {
	for _, handler := range core.currentHandlers {
		handler.Configuration()
	}
}
