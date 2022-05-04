package messaging

type EventDispatcher interface {
	Dispatch(event Event) error
	Dispose()
}
