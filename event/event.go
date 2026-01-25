package event

type Handler func(any)

type EventBus struct {
	listeners map[string][]Handler
}

func NewEventBus() *EventBus {
	return &EventBus{listeners: make(map[string][]Handler)}
}

func (b *EventBus) On(event string, h Handler) {
	b.listeners[event] = append(b.listeners[event], h)
}

func (b *EventBus) Emit(event string, data any) {
	for _, h := range b.listeners[event] {
		go h(data) // async like Node
	}
}
