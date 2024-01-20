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

func (e *EventDispatcher) Register(name string, handler EventHandlerInterface) error {
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

func (e *EventDispatcher) Has(name string, handler EventHandlerInterface) bool {
	if _, ok := e.handlers[name]; ok {
		for _, h := range e.handlers[name] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (e *EventDispatcher) Clear() {
	e.handlers = make(map[string][]EventHandlerInterface)
}
