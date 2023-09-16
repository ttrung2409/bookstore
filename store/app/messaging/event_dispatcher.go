package messaging

import "store/app/domain"

type EventDispatcher interface {
	Dispatch(topic string, key string, events ...domain.Event) error
}
