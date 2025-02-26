package eventdriven

import (
	"sync"
)

// 发布订阅模式
type EventBus struct {
	subscribers sync.Map
}

func (p *EventBus) Subscribe(eventName string, handler func(any)) {
	p.subscribers.Store(eventName, handler)
}

func (p *EventBus) Publish(eventName string, data any) {
	if handler, ok := p.subscribers.Load(eventName); ok {
		handler.(func(any))(data)
	}
}
