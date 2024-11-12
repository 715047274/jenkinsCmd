package application

import (
	"sync"
)

type eventDispatcherImpl struct {
	handlers map[string][]func(interface{})
	mu       sync.RWMutex
}

// NewEventDispatcher creates a new instance of EventDispatcher.
func NewEventDispatcher() EventDispatcher {
	return &eventDispatcherImpl{
		handlers: make(map[string][]func(interface{})),
	}
}

func (d *eventDispatcherImpl) Subscribe(eventType string, handler func(interface{})) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *eventDispatcherImpl) Dispatch(eventType string, event interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if handlers, found := d.handlers[eventType]; found {
		for _, handler := range handlers {
			go handler(event) // Run each handler concurrently
		}
	}
}
