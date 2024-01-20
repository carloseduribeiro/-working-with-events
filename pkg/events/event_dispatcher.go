package events

import (
	"errors"
	"sync"
)

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

func (e *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		wg := new(sync.WaitGroup)
		for _, handler := range handlers {
			wg.Add(1)
			go func(handler EventHandlerInterface) {
				defer wg.Done()
				handler.Handle(event)
			}(handler)
		}
		wg.Wait()
	}
	return nil
}

func (e *EventDispatcher) Remove(name string, handler EventHandlerInterface) {
	if _, ok := e.handlers[name]; ok {
		for i, h := range e.handlers[name] {
			if h == handler {
				e.handlers[name] = append(e.handlers[name][:i], e.handlers[name][i+1:]...)
			}
		}
	}
}

func (e *EventDispatcher) Clear() {
	e.handlers = make(map[string][]EventHandlerInterface)
}
