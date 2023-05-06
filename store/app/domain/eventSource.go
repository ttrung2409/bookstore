package domain

type EventSource interface {
	PendingEvents() []Event
}

type eventSource struct {
	pendingEvents []Event
}

func (es *eventSource) PendingEvents() []Event {
	cloned := make([]Event, len(es.pendingEvents))
	copy(cloned, es.pendingEvents)

	return cloned
}
