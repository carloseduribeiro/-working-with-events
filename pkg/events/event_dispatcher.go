package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e EventDispatcher) Register(name string, handler EventHandlerInterface) error {
	if _, ok := e.handlers[name]; ok {
		for _, h := range e.handlers[name] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	e.handlers[name] = append(e.handlers[name], handler)
	return nil
}
