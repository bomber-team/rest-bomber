package handlers

type IHandlerTopic interface {
	Configuration(signal chan int) error
}
