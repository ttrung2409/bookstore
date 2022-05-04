package kafka

import (
	"context"
	"encoding/json"
	"store/app/messaging"
)

type kafkaEvent struct {
	event messaging.Event
}

func (e *kafkaEvent) Key() string {
	return e.event.Key()
}

func (e *kafkaEvent) Type() string {
	return e.event.Type()
}

func (e *kafkaEvent) Value() []byte {
	serialized, _ := json.Marshal(e.event)

	return serialized
}

type eventDispatcher struct {
	producers map[string]Producer
}

func (d *eventDispatcher) Dispatch(event messaging.Event) error {
	if d.producers == nil {
		d.producers = map[string]Producer{}
	}

	topic := event.Type()
	producer, ok := d.producers[topic]
	if !ok {
		producer = NewProducer(topic)
		d.producers[topic] = producer
	}

	return producer.Send(context.Background(), &kafkaEvent{event})
}

func (d *eventDispatcher) Dispose() {
	if d.producers == nil {
		return
	}

	for _, producer := range d.producers {
		producer.Dispose()
	}
}
