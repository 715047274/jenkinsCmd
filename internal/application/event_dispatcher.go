package application

type EventDispatcher interface {
	Subscribe(eventType string, handler func(interface{}))
	Dispatch(eventType string, event interface{})
}
